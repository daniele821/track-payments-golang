const std = @import("std");

pub fn runCmd(allocator: std.mem.Allocator, cmd: []const []const u8) !void {
    var child = std.process.Child.init(cmd, allocator);
    try child.spawn();
    _ = try child.wait();
}
