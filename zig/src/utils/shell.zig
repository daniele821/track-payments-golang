const std = @import("std");

const CmdOutput = struct {
    stdout: []const u8 = "",
    stderr: []const u8 = "",
    allocator: std.mem.Allocator,

    pub fn deinit(self: @This()) void {
        self.allocator.free(self.stdout);
        self.allocator.free(self.stderr);
    }
};

pub fn runCmd(allocator: std.mem.Allocator, cmd: []const []const u8) !CmdOutput {
    // init
    var cmdOutput = CmdOutput{ .allocator = allocator };
    errdefer cmdOutput.deinit();
    var child = std.process.Child.init(cmd, allocator);
    child.stdout_behavior = .Pipe;
    child.stderr_behavior = .Pipe;

    // run command
    try child.spawn();
    cmdOutput.stdout = try child.stdout.?.readToEndAlloc(allocator, std.math.maxInt(usize));
    cmdOutput.stderr = try child.stderr.?.readToEndAlloc(allocator, std.math.maxInt(usize));
    const result = try child.wait();

    // check command result
    if (result == .Exited and result.Exited == 0) {
        return cmdOutput;
    }
    return error.CommandFailed;
}

test "successful command" {
    const res = try runCmd(std.testing.allocator, ([_][]const u8{ "echo", "CMD OUTPUT" })[0..]);
    res.deinit();
}
