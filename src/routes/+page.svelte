<div class="w-5/6 mx-auto py-5">
    <div class="w-full shadow-2xl subpixel-antialiased rounded h-64 bg-black border-black mx-auto">
        <div class="flex items-center h-6 rounded-t bg-gray-100 border-b border-gray-500 text-center text-black"
             id="headerTerminal">
            <div class="flex ml-2 items-center text-center border-red-900 bg-red-500 shadow-inner rounded-full w-3 h-3"
                 id="closebtn">
            </div>
            <div class="ml-2 border-yellow-900 bg-yellow-500 shadow-inner rounded-full w-3 h-3" id="minbtn">
            </div>
            <div class="ml-2 border-green-900 bg-green-500 shadow-inner rounded-full w-3 h-3" id="maxbtn">
            </div>
            <div class="mx-auto pr-16" id="terminaltitle">
                <p class="text-center text-sm">Praxis Dr. Lenz</p>
            </div>
        </div>
        <div class="pl-1 pt-1 h-auto text-green-200 justify-items-start font-mono text-xs leading-loose truncate tracking-normal bg-black" id="console">
            <p class="pb-1">--Was ist hier passiert?</p>
            <p class="pb-1 whitespace-normal">Diese Seite befindet sich gerade im Umbau!</p>
            <div class="text-justify inline-block">
                <p class="typed truncate text-clip h-5 whitespace-nowrap border-r-2 border-r-orange-400"> Sieh dir das Projekt an:
                    <a class="underline" href="https://github.com/domlen2003/praxislenz.info">Github</a>
                </p>
            </div>
        </div>
    </div>
</div>
<style>
    .typed {
        animation: typing 3s steps(31, end),
        blink-caret .75s step-end infinite;
    }

    /* The typing effect */
    @keyframes typing {
        from {
            width: 0
        }
        to {
            width: 100%
        }
    }

    /* The typewriter cursor effect */
    @keyframes blink-caret {
        from, to {
            border-color: transparent
        }
        50% {
            border-color: orange;
        }
    }
</style>