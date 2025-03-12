const std = @import("std");

const keyBytes = 256 / 8;

// pub fn encrypt(key: [keyBytes]u8) void {
//     std.crypto.core.aes.Aes256.initEnc(key).encryptWide(comptime count: usize, dst: *[16*count]u8, src: *const [16*count]u8)
// }
