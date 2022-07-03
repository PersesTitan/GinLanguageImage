package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var javaScript = "<svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><title>JavaScript</title><path d=\"M0 0h24v24H0V0zm22.034 18.276c-.175-1.095-.888-2.015-3.003-2.873-.736-.345-1.554-.585-1.797-1.14-.091-.33-.105-.51-.046-.705.15-.646.915-.84 1.515-.66.39.12.75.42.976.9 1.034-.676 1.034-.676 1.755-1.125-.27-.42-.404-.601-.586-.78-.63-.705-1.469-1.065-2.834-1.034l-.705.089c-.676.165-1.32.525-1.71 1.005-1.14 1.291-.811 3.541.569 4.471 1.365 1.02 3.361 1.244 3.616 2.205.24 1.17-.87 1.545-1.966 1.41-.811-.18-1.26-.586-1.755-1.336l-1.83 1.051c.21.48.45.689.81 1.109 1.74 1.756 6.09 1.666 6.871-1.004.029-.09.24-.705.074-1.65l.046.067zm-8.983-7.245h-2.248c0 1.938-.009 3.864-.009 5.805 0 1.232.063 2.363-.138 2.711-.33.689-1.18.601-1.566.48-.396-.196-.597-.466-.83-.855-.063-.105-.11-.196-.127-.196l-1.825 1.125c.305.63.75 1.172 1.324 1.517.855.51 2.004.675 3.207.405.783-.226 1.458-.691 1.811-1.411.51-.93.402-2.07.397-3.346.012-2.054 0-4.109 0-6.179l.004-.056z\"/></svg>"
var golang = "<svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><title>Go</title><path d=\"M1.811 10.231c-.047 0-.058-.023-.035-.059l.246-.315c.023-.035.081-.058.128-.058h4.172c.046 0 .058.035.035.07l-.199.303c-.023.036-.082.07-.117.07zM.047 11.306c-.047 0-.059-.023-.035-.058l.245-.316c.023-.035.082-.058.129-.058h5.328c.047 0 .07.035.058.07l-.093.28c-.012.047-.058.07-.105.07zm2.828 1.075c-.047 0-.059-.035-.035-.07l.163-.292c.023-.035.07-.07.117-.07h2.337c.047 0 .07.035.07.082l-.023.28c0 .047-.047.082-.082.082zm12.129-2.36c-.736.187-1.239.327-1.963.514-.176.046-.187.058-.34-.117-.174-.199-.303-.327-.548-.444-.737-.362-1.45-.257-2.115.175-.795.514-1.204 1.274-1.192 2.22.011.935.654 1.706 1.577 1.835.795.105 1.46-.175 1.987-.77.105-.13.198-.27.315-.434H10.47c-.245 0-.304-.152-.222-.35.152-.362.432-.97.596-1.274a.315.315 0 01.292-.187h4.253c-.023.316-.023.631-.07.947a4.983 4.983 0 01-.958 2.29c-.841 1.11-1.94 1.8-3.33 1.986-1.145.152-2.209-.07-3.143-.77-.865-.655-1.356-1.52-1.484-2.595-.152-1.274.222-2.419.993-3.424.83-1.086 1.928-1.776 3.272-2.02 1.098-.2 2.15-.07 3.096.571.62.41 1.063.97 1.356 1.648.07.105.023.164-.117.2m3.868 6.461c-1.064-.024-2.034-.328-2.852-1.029a3.665 3.665 0 01-1.262-2.255c-.21-1.32.152-2.489.947-3.529.853-1.122 1.881-1.706 3.272-1.95 1.192-.21 2.314-.095 3.33.595.923.63 1.496 1.484 1.648 2.605.198 1.578-.257 2.863-1.344 3.962-.771.783-1.718 1.273-2.805 1.495-.315.06-.63.07-.934.106zm2.78-4.72c-.011-.153-.011-.27-.034-.387-.21-1.157-1.274-1.81-2.384-1.554-1.087.245-1.788.935-2.045 2.033-.21.912.234 1.835 1.075 2.21.643.28 1.285.244 1.905-.07.923-.48 1.425-1.228 1.484-2.233z\"/></svg>"
var python = "<svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><title>Python</title><path d=\"M14.25.18l.9.2.73.26.59.3.45.32.34.34.25.34.16.33.1.3.04.26.02.2-.01.13V8.5l-.05.63-.13.55-.21.46-.26.38-.3.31-.33.25-.35.19-.35.14-.33.1-.3.07-.26.04-.21.02H8.77l-.69.05-.59.14-.5.22-.41.27-.33.32-.27.35-.2.36-.15.37-.1.35-.07.32-.04.27-.02.21v3.06H3.17l-.21-.03-.28-.07-.32-.12-.35-.18-.36-.26-.36-.36-.35-.46-.32-.59-.28-.73-.21-.88-.14-1.05-.05-1.23.06-1.22.16-1.04.24-.87.32-.71.36-.57.4-.44.42-.33.42-.24.4-.16.36-.1.32-.05.24-.01h.16l.06.01h8.16v-.83H6.18l-.01-2.75-.02-.37.05-.34.11-.31.17-.28.25-.26.31-.23.38-.2.44-.18.51-.15.58-.12.64-.1.71-.06.77-.04.84-.02 1.27.05zm-6.3 1.98l-.23.33-.08.41.08.41.23.34.33.22.41.09.41-.09.33-.22.23-.34.08-.41-.08-.41-.23-.33-.33-.22-.41-.09-.41.09zm13.09 3.95l.28.06.32.12.35.18.36.27.36.35.35.47.32.59.28.73.21.88.14 1.04.05 1.23-.06 1.23-.16 1.04-.24.86-.32.71-.36.57-.4.45-.42.33-.42.24-.4.16-.36.09-.32.05-.24.02-.16-.01h-8.22v.82h5.84l.01 2.76.02.36-.05.34-.11.31-.17.29-.25.25-.31.24-.38.2-.44.17-.51.15-.58.13-.64.09-.71.07-.77.04-.84.01-1.27-.04-1.07-.14-.9-.2-.73-.25-.59-.3-.45-.33-.34-.34-.25-.34-.16-.33-.1-.3-.04-.25-.02-.2.01-.13v-5.34l.05-.64.13-.54.21-.46.26-.38.3-.32.33-.24.35-.2.35-.14.33-.1.3-.06.26-.04.21-.02.13-.01h5.84l.69-.05.59-.14.5-.21.41-.28.33-.32.27-.35.2-.36.15-.36.1-.35.07-.32.04-.28.02-.21V6.07h2.09l.14.01zm-6.47 14.25l-.23.33-.08.41.08.41.23.33.33.23.41.08.41-.08.33-.23.23-.33.08-.41-.08-.41-.23-.33-.33-.23-.41-.08-.41.08z\"/></svg>"
var csharp = "<svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><title>C Sharp</title><path d=\"M12 0A12 12 0 000 12a12 12 0 0012 12 12 12 0 0012-12A12 12 0 0012 0zM9.426 7.12a5.55 5.55 0 011.985.38v1.181a4.5 4.5 0 00-2.25-.566 3.439 3.439 0 00-2.625 1.087 4.099 4.099 0 00-1.012 2.906 3.9 3.9 0 00.945 2.754 3.217 3.217 0 002.482 1.023 4.657 4.657 0 002.464-.634l-.004 1.08a5.543 5.543 0 01-2.625.555 4.211 4.211 0 01-3.228-1.297 4.793 4.793 0 01-1.212-3.409 5.021 5.021 0 011.365-3.663 4.631 4.631 0 013.473-1.392 5.55 5.55 0 01.12-.004 5.55 5.55 0 01.122 0zm5.863.155h.836l-.555 2.652h1.661l.567-2.652h.81l-.555 2.652 1.732-.004-.15.697H17.91l-.412 1.98h1.852l-.176.698h-1.816l-.58 2.625h-.83l.567-2.625h-1.65l-.555 2.625h-.81l.555-2.625h-1.74l.131-.698h1.748l.401-1.976h-1.826l.138-.697h1.826zm.142 3.345L15 12.6h1.673l.423-1.98z\"/></svg>"
var c = "<svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><title>C</title><path d=\"M16.5921 9.1962s-.354-3.298-3.627-3.39c-3.2741-.09-4.9552 2.474-4.9552 6.14 0 3.6651 1.858 6.5972 5.0451 6.5972 3.184 0 3.5381-3.665 3.5381-3.665l6.1041.365s.36 3.31-2.196 5.836c-2.552 2.5241-5.6901 2.9371-7.8762 2.9201-2.19-.017-5.2261.034-8.1602-2.97-2.938-3.0101-3.436-5.9302-3.436-8.8002 0-2.8701.556-6.6702 4.047-9.5502C7.444.72 9.849 0 12.254 0c10.0422 0 10.7172 9.2602 10.7172 9.2602z\"/></svg>"
var cpp = "<svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><title>C++</title><path d=\"M22.394 6c-.167-.29-.398-.543-.652-.69L12.926.22c-.509-.294-1.34-.294-1.848 0L2.26 5.31c-.508.293-.923 1.013-.923 1.6v10.18c0 .294.104.62.271.91.167.29.398.543.652.69l8.816 5.09c.508.293 1.34.293 1.848 0l8.816-5.09c.254-.147.485-.4.652-.69.167-.29.27-.616.27-.91V6.91c.003-.294-.1-.62-.268-.91zM12 19.11c-3.92 0-7.109-3.19-7.109-7.11 0-3.92 3.19-7.11 7.11-7.11a7.133 7.133 0 016.156 3.553l-3.076 1.78a3.567 3.567 0 00-3.08-1.78A3.56 3.56 0 008.444 12 3.56 3.56 0 0012 15.555a3.57 3.57 0 003.08-1.778l3.078 1.78A7.135 7.135 0 0112 19.11zm7.11-6.715h-.79v.79h-.79v-.79h-.79v-.79h.79v-.79h.79v.79h.79zm2.962 0h-.79v.79h-.79v-.79h-.79v-.79h.79v-.79h.79v.79h.79z\"/></svg>"
var groovy = "<svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><title>Apache Groovy</title><path d=\"M11.997 6.012S10.315 8.8 9.516 10.155c-.155.058-.172.041-.341.207-.41-.47-.897-.041-1.028.22-.057-.566-.151-.567-.279-.694.074-.496.316-1.305-.241-1.884-1.078-.727-2.326 1.05-3.021 1.982l-.375.622c-1.546-.032-2.763.008-4.231-.021 1.79.67 1.864.686 4.026 1.506 0 .066.161.372.34.552.147.15.308.234.389.284-.106.054-.32.138-.385.258-.292.546.139.672.418 1.107.315.568.382.944 1.126.625.254-.11.562-.148.758-.21-.693 1.094-.87 1.392-2.083 3.274l.012.004c4.85-1.893 4.974-1.942 7.373-2.89 3.448 1.338 3.646 1.448 7.432 2.891-.529-.826-.89-1.343-1.274-1.995.151-.013.483-.046.777-.233.213-.135.463-.288.688-.574.443-.565.551-1.277.39-2.166-.078-.423-.235-.834-.213-.85 2.061-.778 2.304-.862 4.226-1.587-2.31.034-2.422.01-4.591.016-.036-.414-.244-.627-.882-.606-.238.039-.389.12-.5.445-.357-.657-.85-.464-1.06-.14-.275-.282-.917-.377-1.24-.17-.238-.112-.514-.112-.757.177-.175-.146-.23-.188-.614-.342-.886-1.497-1.622-2.692-2.36-3.951zm.012.802c.35.535 1.552 2.61 1.849 3.074-.337.023-.668.202-.918.562-.217-.224-.47-.445-.917-.463-.544-.093-.834.148-1.2.568-.108-.365-.53-.45-.896-.28.327-.519 1.872-3.122 2.082-3.46zM7.45 9.128c-.05 1.434-1.068 2.712-1.798 2.245-.551-.449.149-1.584.59-1.985-.033.307.246.498.023.77-.446.543-.27.936-.078.996.513.162 1.004-1.227 1.004-2.201 0-.625-.366-.613-1.086.136-.983 1.022-1.513 2.012-1.16 2.69.197.38.485.651.959.594.925-.11 1.483-1.254 1.543-1.988.148-.003.109.01.148-.02 0 .129.177.755.317 1.166.183.702.964 2.11-1.369 2.658-.44.11-.614.148-1.05.32-.213-.443-.263-.585-.697-1.013.588-.205.593-.185.972-.317 1.467-.51 1.908-.947 1.857-1.57 0 0 .018-.32-.185-.588a2.613 2.613 0 0 1-.293.645c-.437.68-1.296 1.101-2.06.833-.417-.146-.596-.466-.596-1.015 0-.703 1.601-2.735 2.387-3.08.555-.165.579.293.571.724zm6.502 1.3c.26.006.543.133.735.34.594.64.529 1.417.163 1.905-.435.581-1.532.324-1.791-.488-.12-.378.095-1.312.475-1.624a.628.628 0 0 1 .418-.132zm-2.113.066a.502.502 0 0 1 .117.017c.503.03.61.313.701.56.231.626.173 1.212-.301 1.691-.711.719-1.54.401-1.536-.567.014-.69.443-1.715 1.02-1.7zm1.868.038c-.383.287-.432 1.023-.08 1.296.138.13.215.22.613.256.273.024.704-.253.725-.527.01-.125-.013-.333-.312-.67-.252-.283-.579-.349-.661-.3-.265.156.021.28.125.383.162.163.2.234.125.282a.447.447 0 0 1-.372.057c-.105-.049-.456-.246-.163-.777zm3.759.003c.167.26.215.316.402.965.24.838.546 1.163.816 1.01.74-.418.148-1.476-.113-1.974.167-.002.134.007.286.005.12.471.086.387.407 1.813.385 1.706.442 2.16-.528 2.926-.446.352-1.103.37-1.667.34l-.636-.095c.438-.287.545-.557.542-1.116 1.278.535 1.959.132 2.23-.526.132-.317.086-.735-.04-1.471.008.6-.005.71-.084 1.007-.158.595-.547.76-.812.34-.102-.163-.345-.702-.42-1.282-.075-.58-.132-1.395-.5-1.736.04-.08.082-.17.117-.206zm-1.247.01c.258.068.572.204.74.52.234.436.388.668.376 1.447-.014.832-.34 1.055-.557 1.086-.278.04-.762.034-1.049-1.598-.095-.541-.268-1.056-.45-1.224.09-.11.097-.096.165-.204.091.1.17.27.298.777.202.808.387.975.745 1.02.558.072.778-.78.318-1.391-.1-.134-.365-.307-.503-.236.008.236.113.162.114.318-.026.185-.053.219-.113.32-.142-.056-.21-.078-.334-.291-.157-.31-.055-.6.25-.544zm-4.597.076c-.263.185-.594.8-.304 1.35.143.205.297.372.638.3.245-.051.671-.34.73-.749.052-.35-.456-1.028-.738-.87-.327.183-.128.314.074.511.185.18.052.289-.077.342-.258.106-.403.003-.467-.203-.065-.205-.01-.38.144-.68zm-2.867.064c.056.172.1.402.218.624.028.023.132 0 .269-.157.086-.1.185-.238.357-.463.104.095.113.166.142.219.073.13.225.12.273.106.168-.167.195-.275.306-.29.01.216.021.35-.257.677a.535.535 0 0 1-.501.172c-.12-.034-.199-.108-.389-.205-.258.04-.19.315-.143.546.12.611.5.855.832.675.116-.062.09-.062.312-.153-.038.388-.06.463.01.896-.541.301-.982.25-1.102-.506-.091-.632-.261-1.4-.415-1.556-.145-.147-.205-.195-.205-.195l.293-.39zm-7.114.082c.753.01 1.602.01 2.506.017-.13.318-.175.54-.193.854-.422-.163-1.877-.684-2.313-.871zm20.723.01c-.997.359-1.715.637-2.677 1.004-.105-.45-.124-.588-.219-.994 1.601-.005 1.628-.002 2.896-.01zm-6.978 2.04c.105.43.253.641.253.641.202.348.454.545.84.645.085.136.115.163.148.236.037.457.01.514-.344.774-.209.204-.218.497-.003.769.231.22.474.298 1.375.064.174.3.418.653.776 1.217-1.206-.455-2.868-1.103-6.43-2.49 0 0-4.169 1.62-6.404 2.491.935-1.474 1.012-1.599 1.677-2.63.225-.089.149-.053.349-.155.459-.245.827-.61 1.028-1.145.368.83.779.925 1.636.655.177-.082.38-.2.424-.518.46.413 1.432.49 2.142-.382.612.717 2.001.785 2.533-.171zm2.157.865s.04.129.064.169c-.101.003-.213 0-.213 0a.905.905 0 0 0 .149-.17z\"/></svg>"
var ruby = "<svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><title>Ruby</title><path d=\"M20.156.083c3.033.525 3.893 2.598 3.829 4.77L24 4.822 22.635 22.71 4.89 23.926h.016C3.433 23.864.15 23.729 0 19.139l1.645-3 2.819 6.586.503 1.172 2.805-9.144-.03.007.016-.03 9.255 2.956-1.396-5.431-.99-3.9 8.82-.569-.615-.51L16.5 2.114 20.159.073l-.003.01zM0 19.089zM5.13 5.073c3.561-3.533 8.157-5.621 9.922-3.84 1.762 1.777-.105 6.105-3.673 9.636-3.563 3.532-8.103 5.734-9.864 3.957-1.766-1.777.045-6.217 3.612-9.75l.003-.003z\"/></svg>"
var kotlin = "<svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><title>Kotlin</title><path d=\"M24 24H0V0h24L12 12Z\"/></svg>"
var java = "<svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><title>OpenJDK</title><path d=\"M11.915 0 11.7.215C9.515 2.4 7.47 6.39 6.046 10.483c-1.064 1.024-3.633 2.81-3.711 3.551-.093.87 1.746 2.611 1.55 3.235-.198.625-1.304 1.408-1.014 1.939.1.188.823.011 1.277-.491a13.389 13.389 0 0 0-.017 2.14c.076.906.27 1.668.643 2.232.372.563.956.911 1.667.911.397 0 .727-.114 1.024-.264.298-.149.571-.33.91-.5.68-.34 1.634-.666 3.53-.604 1.903.062 2.872.39 3.559.704.687.314 1.15.664 1.925.664.767 0 1.395-.336 1.807-.9.412-.563.631-1.33.72-2.24.06-.623.055-1.32 0-2.066.454.45 1.117.604 1.213.424.29-.53-.816-1.314-1.013-1.937-.198-.624 1.642-2.366 1.549-3.236-.08-.748-2.707-2.568-3.748-3.586C16.428 6.374 14.308 2.394 12.13.215zm.175 6.038a2.95 2.95 0 0 1 2.943 2.942 2.95 2.95 0 0 1-2.943 2.943A2.95 2.95 0 0 1 9.148 8.98a2.95 2.95 0 0 1 2.942-2.942zM8.685 7.983a3.515 3.515 0 0 0-.145.997c0 1.951 1.6 3.55 3.55 3.55 1.95 0 3.55-1.598 3.55-3.55 0-.329-.046-.648-.132-.951.334.095.64.208.915.336a42.699 42.699 0 0 1 2.042 5.829c.678 2.545 1.01 4.92.846 6.607-.082.844-.29 1.51-.606 1.94-.315.431-.713.651-1.315.651-.593 0-.932-.27-1.673-.61-.741-.338-1.825-.694-3.792-.758-1.974-.064-3.073.293-3.821.669-.375.188-.659.373-.911.5s-.466.2-.752.2c-.53 0-.876-.209-1.16-.64-.285-.43-.474-1.101-.545-1.948-.141-1.693.176-4.069.823-6.614a43.155 43.155 0 0 1 1.934-5.783c.348-.167.749-.31 1.192-.425zm-3.382 4.362a.216.216 0 0 1 .13.031c-.166.56-.323 1.116-.463 1.665a33.849 33.849 0 0 0-.547 2.555 3.9 3.9 0 0 0-.2-.39c-.58-1.012-.914-1.642-1.16-2.08.315-.24 1.679-1.755 2.24-1.781zm13.394.01c.562.027 1.926 1.543 2.24 1.783-.246.438-.58 1.068-1.16 2.08a4.428 4.428 0 0 0-.163.309 32.354 32.354 0 0 0-.562-2.49 40.579 40.579 0 0 0-.482-1.652.216.216 0 0 1 .127-.03z\"/></svg>"
var typeScript = "<svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><title>TypeScript</title><path d=\"M1.125 0C.502 0 0 .502 0 1.125v21.75C0 23.498.502 24 1.125 24h21.75c.623 0 1.125-.502 1.125-1.125V1.125C24 .502 23.498 0 22.875 0zm17.363 9.75c.612 0 1.154.037 1.627.111a6.38 6.38 0 0 1 1.306.34v2.458a3.95 3.95 0 0 0-.643-.361 5.093 5.093 0 0 0-.717-.26 5.453 5.453 0 0 0-1.426-.2c-.3 0-.573.028-.819.086a2.1 2.1 0 0 0-.623.242c-.17.104-.3.229-.393.374a.888.888 0 0 0-.14.49c0 .196.053.373.156.529.104.156.252.304.443.444s.423.276.696.41c.273.135.582.274.926.416.47.197.892.407 1.266.628.374.222.695.473.963.753.268.279.472.598.614.957.142.359.214.776.214 1.253 0 .657-.125 1.21-.373 1.656a3.033 3.033 0 0 1-1.012 1.085 4.38 4.38 0 0 1-1.487.596c-.566.12-1.163.18-1.79.18a9.916 9.916 0 0 1-1.84-.164 5.544 5.544 0 0 1-1.512-.493v-2.63a5.033 5.033 0 0 0 3.237 1.2c.333 0 .624-.03.872-.09.249-.06.456-.144.623-.25.166-.108.29-.234.373-.38a1.023 1.023 0 0 0-.074-1.089 2.12 2.12 0 0 0-.537-.5 5.597 5.597 0 0 0-.807-.444 27.72 27.72 0 0 0-1.007-.436c-.918-.383-1.602-.852-2.053-1.405-.45-.553-.676-1.222-.676-2.005 0-.614.123-1.141.369-1.582.246-.441.58-.804 1.004-1.089a4.494 4.494 0 0 1 1.47-.629 7.536 7.536 0 0 1 1.77-.201zm-15.113.188h9.563v2.166H9.506v9.646H6.789v-9.646H3.375z\"/></svg>"

func main() {
	r := gin.Default()

	total := javaScript + "<br>"
	total += golang + "<br>"
	total += python + "<br>"
	total += csharp + "<br>"
	total += c + "<br>"
	total += cpp + "<br>"
	total += groovy + "<br>"
	total += ruby + "<br>"
	total += kotlin + "<br>"
	total += java + "<br>"
	total += typeScript + "<br>"

	r.POST("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"images": total,
		})
	})

	_ = r.Run()
}
