const std = @import("std");
const math = @import("std").math;
const fs = std.fs;

const ArrayList = std.ArrayList;

fn part1(first_col: []i32, second_col: []i32) i32 {
    const first_sorted = first_col;
    const second_sorted = second_col;

    std.mem.sort(i32, first_sorted, {}, comptime std.sort.asc(i32));
    std.mem.sort(i32, second_sorted, {}, comptime std.sort.asc(i32));

    var sum: u32 = 0.0;
    for (first_sorted, 0..) |value, index| {
        sum += @abs(value - second_sorted[index]);
    }

    return @intCast(sum);
}

fn part2(first_col: []const i32, second_col: []const i32) i32 {
    const allocator = std.heap.page_allocator;

    var counts: std.AutoHashMap(i32, i32) = std.AutoHashMap(i32, i32).init(allocator);
    for (second_col) |value| {
        const current_count = counts.get(value) orelse 0;
        counts.put(value, current_count + 1) catch {
            return 0;
        };
    }

    var sum: i32 = 0;
    for (first_col) |value| {
        // Use .get() with a default value of 0
        const count = counts.get(value) orelse 0;
        if (count > 0) {
            sum += value * count;
            counts.put(value, count - 1) catch {
                return 0;
            };
        }
    }

    return sum;
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;

    const args = try std.process.argsAlloc(allocator);
    defer std.process.argsFree(std.heap.page_allocator, args);

    if (args.len < 2) {
        std.debug.panic("Please provide a filename\n", .{});
    }

    const filename = args[1];
    const file = try fs.cwd().openFile(filename, .{});
    defer file.close();

    var first_col = ArrayList(i32).init(allocator);
    var second_col = ArrayList(i32).init(allocator);

    var reader = file.reader();
    var line_buf: [4096]u8 = undefined;

    while (try reader.readUntilDelimiterOrEof(&line_buf, '\n')) |line| {
        var tokens = std.mem.split(u8, line, "   ");

        const first = tokens.next().?;
        const second = tokens.next().?;

        const first_num = try std.fmt.parseInt(i32, first, 10);
        const second_num = try std.fmt.parseInt(i32, second, 10);

        try first_col.append(first_num);
        try second_col.append(second_num);
    }

    const answer1 = part1(first_col.items, second_col.items);
    std.debug.print("{}\n", .{answer1});

    const answer2 = part2(first_col.items, second_col.items);
    std.debug.print("{}\n", .{answer2});
}
