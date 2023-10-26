<script lang="ts">
    import { onMount } from 'svelte'

    type Point = {
        x: number
        y: number
    }

    type Shadow = {
        despawnRadius: number
        direction: number
        location: Point
        maxSize: number
        opacity: number
        origin: Point
        color: string
        speed: number
        size: number
    }

    const blur = 50
    const amount = 12
    const max = 300
    let canvas: HTMLCanvasElement | undefined
    let c: null | CanvasRenderingContext2D
    let shadows: Shadow[] = []
    let isStopped = false
    let iterations = 0
    let start = 0
    let w = 0
    let h = 0

    const animate = () => {
        if (isStopped) return
        requestAnimationFrame(animate)

        w = innerWidth-10
        if (shadows.length < amount) spawn()
        if (shadows.reduce((acc, shadow) => acc + shadow.opacity, 0) < amount/2 - 0.1) spawn()

        if (iterations < 10) iterations++
        switch (iterations) {
            case 5: start = new Date().getTime(); break
            case 10: isStopped = new Date().getTime() - start > max; break
        }

        c!.clearRect(0, 0, w, h)
        for (let i = 0; i < shadows.length; i++) {
            const shadow = shadows[i]
            const { x, y } = shadow.location
            const { x: ox, y: oy } = shadow.origin
            const { despawnRadius, direction, speed, size } = shadow

            c!.beginPath()
            c!.arc(x, y, size, 0, Math.PI * 2, false)
            c!.fillStyle = `rgba(${shadow.color}, ${shadow.opacity})`
            c!.fill()

            if (shadow.size <= shadow.maxSize) {
                shadow.size += 0.5
            }
            const distance = Math.sqrt(Math.pow(x - ox, 2) + Math.pow(y - oy, 2))
            shadow.opacity = Math.max(0, 0.5 - distance / despawnRadius)
            shadow.location = {
                x: x + Math.cos(direction) * speed,
                y: y + Math.sin(direction) * speed
            }
            if (Math.abs(x - ox) > despawnRadius || Math.abs(y - oy) > despawnRadius) shadows.splice(i, 1)
        }
    }

    const spawn = (init = false) => {
        const maxAttempts = 100
        const size = Math.random() * 130 + 170
        for (let i = 0; i < maxAttempts; i++) {
            const x = Math.random() * w
            const y = Math.random() * h / 4
            let skip = false
            for (let j = 0; j < shadows.length; j++) {
                const shadow = shadows[j]
                const { x: sx, y: sy } = shadow.location
                const { despawnRadius } = shadow
                const distance = Math.sqrt(Math.pow(x - sx, 2) + Math.pow(y - sy, 2))
                skip = distance < despawnRadius
            }
            if (skip) continue
            const location = { y, x }
            let color = '168,85,247'
            if (x > w / 2) color = '99,102,241'
            shadows.push({
                despawnRadius: Math.random() * size * 2 + 300,
                speed: Math.random() * 0.5 + 5 / size,
                direction: Math.random() * 180,
                size: init ? size : 25,
                origin: location,
                maxSize: size,
                opacity: 0.5,
                location,
                color,
            })
            break
        }
    }

    $: if (canvas) {
        c = canvas.getContext('2d')
        animate()
    }

    onMount(() => {
        w = innerWidth
        h = innerHeight
        for (let i = 0; i < amount; i++) spawn(true)
    })
</script>

<div class="overflow-hidden bg-transparent absolute" style="width: {w}px; height: {h}px">
    <canvas style="opacity: .7; transform: translateX(-{blur}px), translateY(-{blur*2}px); filter: blur({blur}px)" bind:this={canvas} width={w} height={h} />
</div>
