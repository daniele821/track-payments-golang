const std = @import("std");

pub fn main() !void {
    std.debug.print("testing\n", .{});

    const strings = ([_][]const u8{ "echo", "CMD OUTPUT" })[0..];

    const res = try @import("./utils/shell.zig").runCmd(std.heap.page_allocator, strings, .{});
    res.deinit();
}
