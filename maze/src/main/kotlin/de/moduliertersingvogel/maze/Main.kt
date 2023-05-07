package de.moduliertersingvogel.maze

import kotlin.collections.HashSet
import java.io.FileWriter
import java.nio.file.Path
import java.util.Scanner

fun main(args: Array<String>) {
    val parser = Parser()
    val content = Scanner(Path.of("neighbors.csv")).use { it.useDelimiter("\\Z").next() }
    var maze = parser.parse(content)
    maze = Generation().generate(maze)

    var toexport = ""
    val used = HashSet<Cell>()
    fun export(cell:Cell) {
        used.add(cell)
        val neighbors = cell.neighbors.filter({it.open}).map({it.to}).filter({!used.contains(it)})
        for(n in neighbors) {
            toexport = toexport + if(cell.id < n.id) "${cell.id} ${n.id}\n" else "${n.id} ${cell.id}\n"
            export(n)
        }
    }
    export(maze)
    
    FileWriter("maze.csv").use {it.write(toexport)}
}

