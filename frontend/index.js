function dial() {
    const conn = new WebSocket(`ws://${location.host}/join?seat=32`)

    conn.addEventListener("close", ev => {
        console.log(`WebSocket Disconnected code: ${ev.code}, reason: ${ev.reason}`, true)
        if (ev.code !== 1001) {
            console.log("Reconnecting in 1s", true)
            setTimeout(dial, 1000)
        }
    })
    conn.addEventListener("open", ev => {
        console.info("websocket connected")
    })

    // This is where we handle messages received.
    conn.addEventListener("message", ev => {
        if (typeof ev.data !== "string") {
            console.error("unexpected message type", typeof ev.data)
            return
        }
        const p = appendLog(ev.data)
        if (expectingMessage) {
            p.scrollIntoView()
            expectingMessage = false
        }
    })
}

dial()

let expectingMessage = false
const submit = document.getElementById("mi-boton")

submit.addEventListener("click", async ev => {
    ev.preventDefault()

    const msg = {
        msg: 1,
        seat: 32
    }
    expectingMessage = true
    try {
        const resp = await fetch("/click", {
            method: "POST",
            body: JSON.stringify(msg),
        })
        if(resp.status !== 202) {
            throw new Error(`Unexpected HTTP Status ${resp.status} ${resp.statusText}`)
        }
    } catch (err) {

    }
})