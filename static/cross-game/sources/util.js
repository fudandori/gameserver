import { ACTIVE_COLOUR, INACTIVE_COLOUR } from "./config.js"

/**
 * calculates the {x, y} position inside the matrix based on the index it occupies and decides the color said position depending on the value it holds. 0 for grey and 1 for transparent
 * @param {int} i 
 * @param {int[][]} matrix 
 * @returns object {x, y, color} where x and y are the coordinates inside the matrix and the designated color for the position
 */
export const getInfo = (i, matrix) => {
    const x = Math.floor(i / 5)
    const y = i - (5 * x)
    return { x, y, color: matrix[x][y] === 0 ? INACTIVE_COLOUR : ACTIVE_COLOUR }
}

/**
 * 
 * @param {PointerEvent} clickEvent 
 * @returns array [x, y] indicating the attributes x and y of the element
 */
export const getBox = (clickEvent) => {
    return [
        parseInt(clickEvent.target.getAttribute('x'), 10),
        parseInt(clickEvent.target.getAttribute('y'), 10)
    ]
}