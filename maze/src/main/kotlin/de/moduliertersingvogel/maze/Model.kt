package de.moduliertersingvogel.maze

import kotlin.collections.HashSet

data class Cell(val id: Int) {
    val neighbors = HashSet<Link>()
}

data class Link(val to: Cell, var open:Boolean=false) {
}
