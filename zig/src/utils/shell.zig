const std = @import("std");

const CmdOutput = struct {
    stdout: ?[]const u8 = null,
    stderr: ?[]const u8 = null,
    allocator: std.mem.Allocator,

    pub fn deinit(self: @This()) void {
        if (self.stdout) |stdout| self.allocator.free(stdout);
        if (self.stderr) |stderr| self.allocator.free(stderr);
    }
};

const CmdOptions = struct {
    stdout: bool = false,
    stderr: bool = false,
};

pub fn runCmd(allocator: std.mem.Allocator, cmd: []const []const u8, opts: CmdOptions) !CmdOutput {
    // init
    var cmdOutput = CmdOutput{ .allocator = allocator };
    errdefer cmdOutput.deinit();
    var child = std.process.Child.init(cmd, allocator);
    if (opts.stdout) child.stdout_behavior = .Pipe;
    if (opts.stderr) child.stderr_behavior = .Pipe;

    // run command
    try child.spawn();
    if (child.stdout) |stdout| {
        cmdOutput.stdout = try stdout.readToEndAlloc(allocator, std.math.maxInt(usize));
    }
    if (child.stderr) |stderr| {
        cmdOutput.stderr = try stderr.readToEndAlloc(allocator, std.math.maxInt(usize));
    }
    const result = try child.wait();

    // check command result
    if (result == .Exited and result.Exited == 0) {
        return cmdOutput;
    }
    return error.CommandFailed;
}

test "successful command" {
    const allocator = std.testing.allocator;
    const cmd = ([_][]const u8{ "echo", "CMD OUTPUT" })[0..];
    const opts = .{ .stderr = false, .stdout = true };
    const res = try runCmd(allocator, cmd, opts);
    defer res.deinit();
    try std.testing.expectEqualSlices(u8, "CMD OUTPUT\n", res.stdout.?);
    try std.testing.expect(res.stderr == null);
}

test "fail command" {
    const allocator = std.testing.allocator;
    const cmd = ([_][]const u8{"adawdawdawdawd"})[0..];
    const opts = .{ .stderr = false, .stdout = true };
    _ = runCmd(allocator, cmd, opts) catch |err| {
        try std.testing.expectEqual(err, error.FileNotFound);
        return;
    };
}
