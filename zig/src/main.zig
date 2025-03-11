const std = @import("std");

pub fn main() !void {
    std.debug.print("testing\n", .{});

    const strings = [_][]const u8{ "echo", "ciao", "" };

    const slice_of_slices: []const []const u8 = strings[0..];
    try @import("./utils/shell.zig").runCmd(std.heap.page_allocator, slice_of_slices);
}
