package de.moduliertersingvogel.dojos.brain

import org.junit.jupiter.api.Test
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Assertions.assertTrue

import io.mockk.mockkStatic
import io.mockk.every
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
        mockkStatic(::moveleft)
        mockkStatic(::incr)
        mockkStatic(::decr)
        mockkStatic(::outp)
        mockkStatic(::inp)
        mockkStatic(::jumpf)
        mockkStatic(::jumpb)
        every {outp(any())} returns Unit
        every {inp(any())} returns Unit

        var test = "><+-.,[]"
        Brain().parse(test)

        verify { moveright(any()) }
        verify { moveleft(any()) }
        verify { incr(any()) }
        verify { decr(any()) }
        verify { outp(any()) }
        verify { inp(any()) }
        verify { jumpf(any(), any()) }
        verify { jumpb(any(), any()) }
    }

    @Test
    fun testmoveright() {
        val state = BrainState()
        assertEquals(0u, state.pointer)
        moveright(state)
        assertEquals(1u, state.pointer)
    }

    @Test
    fun testmoveleft() {
        val state = BrainState()
        state.pointer = 1u
        moveleft(state)
        assertEquals(0u, state.pointer)
    }

    @Test
    fun testincr() {
        val state = BrainState()
        state.tape[0] = 0u
        incr(state)
        assertEquals(1u, state.tape[0])
    }

    @Test
    fun testdecr() {
        val state = BrainState()
        state.tape[0] = 1u
        decr(state)
        assertEquals(0u, state.tape[0])
    }

    @Test
    fun testoutput() {
        mockkStatic(::output)
        every { output(any()) } returns Unit
        val state = BrainState()
        outp(state)

        verify { output(any()) }
    }

    @Test
    fun testinput() {
        mockkStatic(::input)
        every { input() } returns 0u
        val state = BrainState()
        inp(state)

        verify { input() }
    }

    @Test
    fun testjumpf() {
        val state = BrainState()
        val programm = "[>>>>]"
        assertEquals(0u, state.pointer)
        jumpf(state, programm)
        assertEquals(6u, state.pointer)
    }

    @Test
    fun testjumpb() {
        val state = BrainState()
        state.tape[0] = 1u
        decr(state)
        assertEquals(0u, state.tape[0])
    }
}
