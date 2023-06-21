package de.moduliertersingvogel.dojos.brain

import kotlin.collections.mapOf

fun moveright(state: BrainState, programm: String = "", index: Int = 0): Int {
    state.pointer = state.pointer + 1u
    return 1
}

fun moveleft(state: BrainState, programm: String = "", index: Int = 0): Int {
    state.pointer = state.pointer - 1u
    return 1
}

fun incr(state: BrainState, programm: String = "", index: Int = 0): Int {
    state.tape[state.pointer.toInt()] = state.tape[state.pointer.toInt()] + 1u
    return 1
}

fun decr(state: BrainState, programm: String = "", index: Int = 0): Int {
    state.tape[state.pointer.toInt()] = state.tape[state.pointer.toInt()] - 1u
    return 1
}

fun outp(state: BrainState, programm: String = "", index: Int = 0): Int {
    output(state.tape[state.pointer.toInt()])
    return 1
}

fun inp(state: BrainState, programm: String = "", index: Int = 0): Int {
    state.tape[state.pointer.toInt()] = input()
    return 1
}

fun jumpf(state: BrainState, programm: String, index: Int = 0): Int {
    if(state.tape[state.pointer.toInt()] == 0u) {
        for(i: Int in 0 .. programm.length - index) {
            if(programm[index + i] == ']') {
                return i
            }
        }
    }
    return 1
}

fun jumpb(state: BrainState, programm: String, index: Int = 0): Int {
    if(state.tape[state.pointer.toInt()] > 0u) {
        for(i: Int in 0 .. index) {
            if(programm[index - i] == '[') {
                return -1 * i
            }
        }
    }
    return 1
}

class BrainState() {
    var pointer = 0u;
    val tape = Array(1024) {_ -> 0u}
}

class Brain() {
    var position = 0

    fun parse(programm: String) {
        val state = BrainState()
        val dict = mapOf('>' to ::moveright,'<' to ::moveleft, '+' to ::incr, '-' to ::decr, ',' to ::outp, '.' to ::inp, '[' to ::jumpf, ']' to ::jumpb)
        var i = 0
        while(i < programm.length) {
            i = i + dict[programm[i]]!!.invoke(state, programm, i)
        }
    }
}
