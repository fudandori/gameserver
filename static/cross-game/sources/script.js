import { calculateURL, generateProps, generateURL, INACTIVE_COLOUR } from "./config.js"
import { getBox, getInfo } from "./util.js"

let matrix = Array(5)
let active = true
const board = document.getElementById('board')

/**
 * initializes the board
 */
const start = () => {
    call(generateURL, generateProps).then(r => {
        matrix = r.matrix

        for (let i = 0; i < 25; i++) {
            const { x, y, color } = getInfo(i, matrix)

            let div = document.createElement('div')
            div.classList.add('box')
            div.setAttribute('x', x)
            div.setAttribute('y', y)
            div.style.backgroundColor = color
            div.addEventListener('click', send)
            board.appendChild(div)
        }
    })
}

/**
 * sends a request to the server for a box switch
 * @param {PointerEvent} event the pointer event to extract the box clicked
 */
const send = (event) => {
    if (!active) return

    const body = JSON.stringify({ matrix, box: getBox(event) })
    const props = { method: "POST", body }

    const children = board.children
    call(calculateURL, props).then(r => {
        matrix = r.matrix
        for (let i = 0; i < 25; i++) {
            children[i].style.backgroundColor = getInfo(i, matrix).color
        }

        checkForWin()
    })
}

/**
 * performs a fetch to the designated URL with the provided properties
 * @param {string} url 
 * @param {{method: string, body: string}} props 
 * @returns promise with the JSON response unwrapped
 */
const call = (url, props) => {
    return fetch(url, props).then(res => { return res.json() })
}

/**
 * checks if there is a winning condition and if so, activates the win mechanism
 */
const checkForWin = () => {

    const children = board.children
    let inactiveFound = false
    for (let i = 0; i < 25 && !inactiveFound; i++) {
        inactiveFound = children[i].style.backgroundColor === INACTIVE_COLOUR
    }

    if (!inactiveFound) {
        document.getElementById('modal').classList.add('anim')
        board.classList.add('win')
        active = false
    }
}

start()

export default { matrix, send }
