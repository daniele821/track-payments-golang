const std = @import("std");

const CmdOutput = struct {
    stdout: []const u8 = "",
    stderr: []const u8 = "",
};

pub fn runCmd(allocator: std.mem.Allocator, cmd: []const []const u8) !CmdOutput {
    var cmdOutput = CmdOutput{};

    var child = std.process.Child.init(cmd, allocator);
    child.stdout_behavior = .Pipe;
    child.stderr_behavior = .Pipe;

    try child.spawn();

    cmdOutput.stdout = try child.stdout.?.readToEndAlloc(allocator, std.math.maxInt(usize));
    cmdOutput.stderr = try child.stderr.?.readToEndAlloc(allocator, std.math.maxInt(usize));

    const result = try child.wait();
    if (result == .Exited and result.Exited == 0) {
        return cmdOutput;
    }
    return error.CommandFailed;
}
