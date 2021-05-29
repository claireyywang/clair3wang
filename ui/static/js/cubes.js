import * as THREE from 'https://cdn.skypack.dev/three@0.128.0/build/three.module.js'
import { OrbitControls } from 'https://cdn.skypack.dev/three@0.128.0/examples/jsm/controls/OrbitControls.js'

function main() {
    // canvas
    const canvas = document.querySelector("#art");

    // renderer
    const renderer = new THREE.WebGLRenderer({canvas});

    // camera
    const fov = 75;
    const aspect = 2;
    const near = 0.1;
    const far = 5;
    const camera = new THREE.PerspectiveCamera( fov, aspect, near, far );
    camera.position.z = 2;

    // scene
    const scene = new THREE.Scene();
    scene.background = new THREE.Color('azure');

     // light
     {
        const color = 0xffffff;
        const intensity = 1;
        const light = new THREE.DirectionalLight(color, intensity);
        light.position.set(-1, 2, 4);
        scene.add(light);
    }

    // box
    const boxWidth = .5;
    const boxHeight = .5;
    const boxDepth = .5;
    const geometry = new THREE.BoxGeometry(boxWidth, boxHeight, boxDepth);

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

    function resizeRendererToDisplaySize(renderer) {
        const canvas = renderer.domElement;
        const pixelRatio = window.devicePixelRatio;
        const width = canvas.clientWidth * pixelRatio | 0;
        const height = canvas.clientHeight * pixelRatio | 0;
        const needResize = canvas.width !== width || canvas.height !== height;
        if (needResize) {
            renderer.setSize(width, height, false);
        }
        return needResize;
    }

    function render(time) {
        time *= 0.001;

        if (resizeRendererToDisplaySize(renderer)) {
            const canvas = renderer.domElement;
            camera.aspect = canvas.clientWidth / canvas.clientHeight;
            camera.updateProjectionMatrix();
        }

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