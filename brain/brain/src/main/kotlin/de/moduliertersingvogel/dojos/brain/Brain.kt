package de.moduliertersingvogel.dojos.brain

import kotlin.collections.mapOf

fun moveright() {
}

fun moveleft() {
}

fun incr() {
}

fun decr() {
}

fun outp() {
}

fun inp() {
}

fun jumpf() {
}

fun jumpb() {
}


class Brain() {
    var position = 0

    fun parse(programm: String) {
        val dict = mapOf('>' to ::moveright,'<' to ::moveleft, '+' to ::incr, '-' to ::decr, ',' to ::outp, '.' to ::inp, ']' to ::jumpf, '[' to ::jumpb)
        for(i in 0 until programm.length) {
            dict[programm[i]]?.invoke()
        }
    }
}
