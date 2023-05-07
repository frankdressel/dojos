package de.moduliertersingvogel.maze

import kotlin.test.assertContentEquals
import kotlin.test.assertEquals

import org.junit.jupiter.api.BeforeAll
import org.junit.jupiter.api.Test

class ParserTest {
    companion object {
        private lateinit var testset: String

        @BeforeAll
        @JvmStatic
        fun setup() {
            testset = """
                1 2
                1 3
                2 3
                2 5
                3 4
            """.trimIndent()
        }
    }

    @Test
    fun `Nodes should be parsed correctly`() {
        val parser = Parser()
        val cell = parser.parse(ParserTest.testset)
        
        assertEquals(1, cell.id)
    }

    @Test
    fun `Links should be established`() {
        val parser = Parser()
        val cell1 = parser.parse(ParserTest.testset)
        assertEquals(setOf(Link(Cell(2)), Link(Cell(3))), cell1.neighbors)

        val cell2 = cell1.neighbors.filter({it.to.id == 2}).first().to
        assertEquals(setOf(Link(Cell(1)), Link(Cell(3)), Link(Cell(5))), cell2.neighbors)

        val cell3 = cell1.neighbors.filter({it.to.id == 3}).first().to
        assertEquals(setOf(Link(Cell(1)), Link(Cell(2)), Link(Cell(4))), cell3.neighbors)
    }
}
