package de.moduliertersingvogel.dojos.brain

import org.junit.jupiter.api.Test
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Assertions.assertTrue

import io.mockk.mockkStatic
import io.mockk.verify

import de.moduliertersingvogel.dojos.brain.moveright

class BrainTest {
    @Test
    fun testinit() {
        Brain()
    }

    @Test
    fun testfunctioncalls() {
        mockkStatic(::moveright)

        var test = "><+-.,[]"
        Brain().parse(test)

        verify { moveright() }
        verify { moveleft() }
        verify { incr() }
        verify { decr() }
        verify { outp() }
        verify { inp() }
        verify { jumpf() }
        verify { jumpb() }
    }
}
