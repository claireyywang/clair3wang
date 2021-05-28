import * as THREE from 'https://cdn.skypack.dev/three@0.128.0'

function main() {
    // canvas
    const canvas = document.querySelector("#art");
    const width = window.innerWidth * 0.75;
    const height = width / 1.26;

    // renderer
    const renderer = new THREE.WebGLRenderer({canvas});
    renderer.setSize(width, height);
    // resize when window size changes
    window.addEventListener("resize", function() {
        renderer.setSize(width, height);
    });

    // camera
    const fov = 75;
    // const aspect = 2;
    const near = 0.1;
    const far = 5;
    const camera = new THREE.PerspectiveCamera( fov, width/height, near, far );
    camera.position.z = 2;

    // scene
    const scene = new THREE.Scene();
    scene.background = new THREE.Color('azure');

    // box
    const boxWidth = .5;
    const boxHeight = .5;
    const boxDepth = .5;
    const geometry = new THREE.BoxGeometry(boxWidth, boxHeight, boxDepth);

    // light
    {
        const color = 0xffffff;
        const intensity = 1;
        const light = new THREE.DirectionalLight(color, intensity);
        light.position.set(-1, 2, 4);
        scene.add(light);
    }

    function makeObject(geometry, color, x) {
        const material = new THREE.MeshPhongMaterial({color});
        const cube = new THREE.Mesh(geometry, material);
        scene.add(cube);
        cube.position.x = x;
        return cube;
    }

    const cubes = [
        makeObject(geometry, 0xcd5c5c, 0),
        makeObject(geometry, 0x228b22, -1),
        makeObject(geometry, 0x4169e1, 1),
    ]

    function render(time) {
        time *= 0.001;

        cubes.forEach((cube, ndx) => {
            const speed = 1 + ndx * .1;
            const rot = time * speed;
            cube.rotation.x = rot;
            cube.rotation.y = rot;
        });

        renderer.render(scene, camera)
        requestAnimationFrame(render);
    }
    requestAnimationFrame(render);
}

main();