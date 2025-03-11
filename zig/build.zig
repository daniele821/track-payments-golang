const std = @import("std");

pub fn build(b: *std.Build) void {
    const exe = b.addExecutable(.{
        .name = "payments",
        .root_source_file = b.path("src/main.zig"),
        .target = b.standardTargetOptions(.{}),
        .optimize = b.standardOptimizeOption(.{}),
    });
    const run_exe = b.addRunArtifact(exe);

    b.installArtifact(exe);

    const run_step = b.step("run", "run the program");
    run_step.dependOn(&run_exe.step);
}
