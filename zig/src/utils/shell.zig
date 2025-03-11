const std = @import("std");

pub fn runCmd(allocator: std.mem.Allocator, cmd: []const []const u8) !void {
    var child = std.process.Child.init(cmd, allocator);
    child.stdout_behavior = .Pipe;
    child.stderr_behavior = .Pipe;

    try child.spawn();

    const stdout = try child.stdout.?.readToEndAlloc(allocator, std.math.maxInt(usize));
    const stderr = try child.stderr.?.readToEndAlloc(allocator, std.math.maxInt(usize));

    const term = try child.wait();

    std.debug.print("{s}{s}{}\n", .{ stdout, stderr, term });
}
