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
        every {outp(any())} returns 1
        every {inp(any())} returns 1

        var test = "++[-><.,]"
        Brain().parse(test)

        verify { moveright(any(), any(), any()) }
        verify { moveleft(any(), any(), any()) }
        verify { incr(any(), any(), any()) }
        verify { decr(any(), any(), any()) }
        verify { outp(any(), any(), any()) }
        verify { inp(any(), any(), any()) }
        verify { jumpf(any(), any(), any()) }
        verify { jumpb(any(), any(), any()) }
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
        state.tape[2] = 0u
        state.pointer = 2u
        incr(state)
        assertEquals(1u, state.tape[2])
    }

    @Test
    fun testdecr() {
        val state = BrainState()
        state.tape[2] = 1u
        state.pointer = 2u
        decr(state)
        assertEquals(0u, state.tape[2])
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
    fun testjumpb() {
        var state = BrainState()
        var programm = "[++]"
        var jumped = jumpb(state, programm, 3)
        assertEquals(1, jumped)

        state = BrainState()
        programm = "[++]"
        state.tape[0] = 1u
        jumped = jumpb(state, programm, 3)
        assertEquals(-3, jumped)
    }

    @Test
    fun testjumpf() {
        var state = BrainState()
        var programm = "[++]"
        var jumped = jumpf(state, programm, 0)
        assertEquals(3, jumped)

        state = BrainState()
        programm = "[++]"
        state.tape[0] = 1u
        jumped = jumpf(state, programm, 3)
        assertEquals(1, jumped)
    }

    @Test
    fun testparse() {
        var programm = "++[>+<-]"
        val parsed = Brain().parse(programm)
        assertEquals(0u, parsed[0])
        assertEquals(2u, parsed[1])
    }
}
