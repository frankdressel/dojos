package de.moduliertersingvogel.maze

import org.junit.jupiter.api.Test
import kotlin.test.assertEquals
import kotlin.test.assertFalse
import kotlin.test.assertTrue

class ModelTest {
    @Test
    fun `Cell should be created without errors`() {
        val cell = Cell(1)
    }

    @Test
    fun `Link should be created without errors`() {
        val link = Link(Cell(1))
    }

    @Test
    fun `hashCode of cell class should be immutable`() {
        val c = Cell(1)
        val h1 = c.hashCode()
        c.neighbors.add(Link(Cell(2)))
        val h2 = c.hashCode()

        assertEquals(h1, h2)
    }

    @Test
    fun `Links can be open or closed`() {
        val link = Link(Cell(1))
        link.open = true
        assertTrue(link.open)
        link.open = false
        assertFalse(link.open)
        link.open = true
        assertTrue(link.open)
    }

    @Test
    fun `Links can be assigned to cells`() {
        val cell = Cell(1)
        cell.neighbors.add(Link(Cell(2)))
        cell.neighbors.add(Link(Cell(3)))

        assertEquals(2, cell.neighbors.size)
    }
}
