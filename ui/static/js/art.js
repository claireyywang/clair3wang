import * as THREE from 'https://cdn.skypack.dev/three@0.128.0'

const container = document.getElementById("art")
const width = window.innerWidth * 0.75;
const height = width / 1.26;

const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera( 75, width / height , 0.1, 1000 );
camera.position.z = 5;

const renderer = new THREE.WebGLRenderer();
renderer.setSize(width, height);
// resize when window size changes
window.addEventListener("resize", function() {
    renderer.setSize(width, height);
});
container.appendChild( renderer.domElement );

// create cube
const geometry = new THREE.BoxGeometry();
const material = new THREE.MeshBasicMaterial( { color: 0x00ff00 } );
const cube = new THREE.Mesh( geometry, material );
scene.add( cube );

const animate = function () {
    requestAnimationFrame( animate );

    cube.rotation.x += 0.01;
    cube.rotation.y += 0.01;

    renderer.render( scene, camera );
};

animate();