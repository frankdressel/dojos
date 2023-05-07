package de.moduliertersingvogel.maze

import kotlin.collections.HashMap

class Parser() {
    fun parse(content: String): Cell {
        val cells = HashMap<Int, Cell>()
        val lines = content.split("\n")
        for(line in lines) {
            val pair = line.split(" ")
            val p0 = pair[0].toInt()
            val p1 = pair[1].toInt()
            val left = if (cells.containsKey(p0)) cells.get(p0) else {cells.put(p0, Cell(p0)); cells.get(p0)}
            val right = if (cells.containsKey(p1)) cells.get(p1) else {cells.put(p1, Cell(p1)); cells.get(p1)}

            val lr = Link(right!!)
            val rl = Link(left!!)
            left.neighbors.add(lr)
            right.neighbors.add(rl)
        }
        return cells.get(1)!!
    }
}
