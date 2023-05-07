package de.moduliertersingvogel.maze

import kotlin.collections.HashMap
import kotlin.collections.HashSet

import org.junit.jupiter.api.BeforeAll
import org.junit.jupiter.api.Test

import kotlin.test.assertEquals
import kotlin.test.assertFalse
import kotlin.test.assertTrue

class GenerationTest {
    companion object {
        private lateinit var rectangular_root: Cell
        private lateinit var triangle_root: Cell

        @BeforeAll
        @JvmStatic
        fun setup() {
            fun setup_rect() {
                val cells = HashMap<Int, Cell>()
                for(i in 1..4) {
                    cells.put(i, Cell(i))
                }

                cells.get(1)!!.neighbors.add(Link(cells.get(2)!!))
                cells.get(1)!!.neighbors.add(Link(cells.get(3)!!))
                cells.get(2)!!.neighbors.add(Link(cells.get(1)!!))
                cells.get(2)!!.neighbors.add(Link(cells.get(4)!!))
                cells.get(3)!!.neighbors.add(Link(cells.get(1)!!))
                cells.get(3)!!.neighbors.add(Link(cells.get(4)!!))
                cells.get(4)!!.neighbors.add(Link(cells.get(2)!!))
                cells.get(4)!!.neighbors.add(Link(cells.get(3)!!))

                this.rectangular_root = cells.get(1)!!
            }

            fun setup_trian() {
                val cells = HashMap<Int, Cell>()
                for(i in 1..6) {
                    cells.put(i, Cell(i))
                }

                cells.get(1)!!.neighbors.add(Link(cells.get(2)!!))
                cells.get(1)!!.neighbors.add(Link(cells.get(6)!!))
                cells.get(2)!!.neighbors.add(Link(cells.get(1)!!))
                cells.get(2)!!.neighbors.add(Link(cells.get(3)!!))
                cells.get(3)!!.neighbors.add(Link(cells.get(2)!!))
                cells.get(3)!!.neighbors.add(Link(cells.get(4)!!))
                cells.get(4)!!.neighbors.add(Link(cells.get(3)!!))
                cells.get(4)!!.neighbors.add(Link(cells.get(5)!!))
                cells.get(5)!!.neighbors.add(Link(cells.get(4)!!))
                cells.get(5)!!.neighbors.add(Link(cells.get(6)!!))
                cells.get(6)!!.neighbors.add(Link(cells.get(5)!!))
                cells.get(6)!!.neighbors.add(Link(cells.get(1)!!))

                this.triangle_root = cells.get(1)!!
            }

            setup_rect()
            setup_trian()
        }
    }

    @Test
    fun `Simple rectangular grid`() {
        val maze = Generation().generate(GenerationTest.rectangular_root)

        assertEquals(1, maze.id)

        val openlinks = maze.neighbors.filter {it.open}
        assertEquals(1, openlinks.size)
    }

    @Test
    fun `Simple triangular grid`() {
        var maze = Generation().generate(GenerationTest.triangle_root)

        assertEquals(1, maze.id)
        assertEquals(1, maze.neighbors.filter({it.open}).size)
        val direction = maze.neighbors.filter({it.open}).get(0).to.id

        val openlinks = maze.neighbors.filter {it.open}
        assertEquals(1, openlinks.size)
        val visited = HashSet<Cell>()
        visited.add(maze)
        while(maze.neighbors.filter({it.open}).filter({!visited.contains(it.to)}).size > 0) {
            maze = maze.neighbors.filter({it.open}).filter({!visited.contains(it.to)}).get(0).to
            if(direction == 2 && maze.id != 6) {
                assertEquals(2, maze.neighbors.filter({it.open}).size)
            }
            if(direction == 6 && maze.id != 2) {
                assertEquals(2, maze.neighbors.filter({it.open}).size)
            }
            visited.add(maze)
        }
        assertEquals(1, maze.neighbors.filter({it.open}).size)
    }
}
