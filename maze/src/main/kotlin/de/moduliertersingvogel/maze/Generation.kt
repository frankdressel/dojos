package de.moduliertersingvogel.maze

import kotlin.collections.HashSet
import java.util.WeakHashMap
import org.apache.logging.log4j.kotlin.logger

class Generation {
    private val _visited = WeakHashMap<Cell, Boolean>()
    val logger = logger("maze")

    var Cell.visited: Boolean
        get() = _visited.containsKey(this)
        set(v) {
            if(v) _visited.put(this, true) else _visited.remove(this)
        }

    fun search(current: Cell) {
        logger.info("Searching with ${current}")
        if(current.visited) {
            logger.info("Ignoring: ${current}")
            return
        }

        current.visited = true
        val neighbors = current.neighbors.filter({!it.open}).shuffled()

        for(next in neighbors) {
            if(!next.open && !next.to.visited) {
                logger.info("Searching neighbor ${next} for cell ${current}")
                next.open = true
                next.to.neighbors.filter({it.to == current}).get(0)!!.open = true
                search(next.to)
            }
        }
    }

    fun generate(root: Cell):Cell {
        this._visited.clear()

        search(root)

        return root
    }
}
