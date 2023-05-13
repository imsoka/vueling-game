function setIdValue() {
    let id_header = document.getElementById("idjugador");
    id_header.innerHTML = id;
}

function getRandomInt(max) {
    return Math.floor(Math.random() * max);
}

let id = getRandomInt(100);
setIdValue();


function dial() {
    const conn = new WebSocket(`ws://${location.host}/join?seat=${id}`)

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

    expectingMessage = true
    try {
        const resp = await fetch("/click", {
            method: "POST",
            body: JSON.stringify(
                {
                    msg: "click",
                    seat: `${id}`
                }),
        })
        if(resp.status !== 202) {
            throw new Error(`Unexpected HTTP Status ${resp.status} ${resp.statusText}`)
        }
    } catch (err) {

    }
})


let contador = 0;
var contando = false;
const boton = document.getElementById("mi-boton");
const contadorSpan = document.getElementById("contador");

boton.addEventListener("click", () => {
    contador++;
    contadorSpan.textContent = contador;
});


// Obtener el elemento del contador y el botón
var contadorElemento = document.getElementById("contador2");
var botonn = document.getElementById("mi-boton");

// Definir el contador inicial
var contador2 = 10;

// Función que se llama cuando se hace clic en el botón
function comenzarContador() {
    // Deshabilitar el botón para evitar clics repetidos
    boton.disabled = false;
    if(!contando) {
        contando = true

        // Crear un intervalo que disminuya el contador cada segundo
        var intervalo = setInterval(function() {
            // Disminuir el contador2
            contador2--;

            // Actualizar el elemento del contador
            contadorElemento.innerHTML = contador2;



            var miElementox = document.getElementById("contador");
            var mostrarsegunda = miElementox.innerHTML;
            var titulo = document.getElementById("titulo");
            titulo.innerHTML = mostrarsegunda;



            // Si el contador llega a cero, mostrar el pop-up y detener el intervalo
            if (contador2 == 0) {
                botonn.disabled = true;
                //alert("Se acabó el tiempo");
                clearInterval(intervalo);
                function mostrarDiv() {
                    var div = document.getElementById("miDiv");
                    div.style.display = "block";
                }
                mostrarDiv();

                // Hacer F5 después de cerrar el pop-up
                //location.reload();
            }
        }, 1000); // 1000 milisegundos = 1 segundo
    } 
}

// Agregar el evento click al botón
boton.addEventListener("click", comenzarContador);
