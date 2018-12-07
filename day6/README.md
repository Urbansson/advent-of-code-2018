


## --- Day 6: Chronal Coordinates ---

The device on your wrist beeps several times, and once again you feel like you&#39;re falling.

&#34;
Situation critical
,&#34; the device announces. &#34;Destination indeterminate. Chronal interference detected. Please specify new target coordinates.&#34;

The device then produces a list of coordinates (your puzzle input). Are they places it thinks are safe or dangerous? It recommends you check manual page 729. The Elves did not give you a manual.

_If they&#39;re dangerous,_ maybe you can minimize the danger by finding the coordinate that gives the largest distance from the other points.

Using only the [Manhattan distance](https://en.wikipedia.org/wiki/Taxicab_geometry), determine the _area_ around each coordinate by counting the number of [integer](https://en.wikipedia.org/wiki/Integer) X,Y locations that are _closest_ to that coordinate (and aren&#39;t _tied in distance_ to any other coordinate).

Your goal is to find the size of the _largest area_ that isn&#39;t infinite. For example, consider the following list of coordinates:

``1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
``

If we name these coordinates `A` through `F`, we can draw them on a grid, putting `0,0` at the top left:

``..........
.A........
..........
........C.
...D......
.....E....
.B........
..........
..........
........F.
``

This view is partial - the actual grid extends infinitely in all directions.  Using the Manhattan distance, each location&#39;s closest coordinate can be determined, shown here in lowercase:

``aaaaa.cccc
a_A_aaa.cccc
aaaddecccc
aadddecc_C_c
..d_D_deeccc
bb.de_E_eecc
b_B_b.eeee..
bbb.eeefff
bbb.eeffff
bbb.ffff_F_f
``

Locations shown as `.` are equally far from two or more coordinates, and so they don&#39;t count as being closest to any.

In this example, the areas of coordinates A, B, C, and F are infinite - while not shown here, their areas extend forever outside the visible grid. However, the areas of coordinates D and E are finite: D is closest to 9 locations, and E is closest to 17 (both including the coordinate&#39;s location itself).  Therefore, in this example, the size of the largest area is _17_.

_What is the size of the largest area_ that isn&#39;t infinite?


