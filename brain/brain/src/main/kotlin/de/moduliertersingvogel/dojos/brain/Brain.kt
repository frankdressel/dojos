package de.moduliertersingvogel.dojos.brain

import kotlin.collections.mapOf

fun moveright(state: BrainState) {
    state.pointer = state.pointer + 1u
}

fun moveleft(state: BrainState) {
    state.pointer = state.pointer - 1u
}

fun incr(state: BrainState) {
    state.tape[state.pointer.toInt()] = state.tape[state.pointer.toInt()] + 1u
}

fun decr(state: BrainState) {
    state.tape[state.pointer.toInt()] = state.tape[state.pointer.toInt()] - 1u
}

fun outp(state: BrainState) {
    output(state.tape[state.pointer.toInt()])
}

fun inp(state: BrainState) {
    state.tape[state.pointer.toInt()] = input()
}

fun jumpf(state: BrainState, programm: String) {
    for(i: Int in state.pointer.toInt() .. programm.length) {
        if(programm[i] == ']') {
            state.pointer = (i + 1).toUInt()
            return
        }
    }
}

fun jumpb(state: BrainState, programm: String) {
    for(i: Int in state.pointer.toInt() - 1 downTo 0) {
        if(programm[i] == '[') {
            state.pointer = i.toUInt()
            return
        }
    }
}

class BrainState() {
    var pointer = 0u;
    val tape = Array(1024) {_ -> 0u}
}

class Brain() {
    var position = 0

    fun parse(programm: String) {
        val state = BrainState()
        val dict = mapOf('>' to ::moveright,'<' to ::moveleft, '+' to ::incr, '-' to ::decr, ',' to ::outp, '.' to ::inp, '[' to {s: BrainState -> jumpf(s, programm)}, ']' to {s: BrainState -> jumpb(s, programm)})
        for(i in 0 until programm.length) {
            dict[programm[i]]?.invoke(state)
        }
    }
}
