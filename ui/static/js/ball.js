import * as THREE from 'https://cdn.skypack.dev/three@0.128.0'
import {OrbitControls} from "https://threejs.org/examples/jsm/controls/OrbitControls.js";

// canvas size
const container = document.getElementById("art")
const width = window.innerWidth * 0.75;
const height = width / 1.26;

// camera
const camera = new THREE.PerspectiveCamera( 45, width / height , 1, 80000 );
camera.position.set(-300, 300, -1000);
camera.lookAt(0, 0, 0);

// light
const ambientLight = new THREE.AmbientLight(0xffffff);
const light = new THREE.DirectionalLight(0xffffff, 0.7);
light.position.set(-800, 900, 300);

// renderer
const renderer = new THREE.WebGLRenderer();
renderer.setSize(width, height);
// resize when window size changes
window.addEventListener("resize", function() {
    renderer.setSize(width, height);
});
container.appendChild( renderer.domElement );

// control
const controls = new OrbitControls(camera, renderer.domElement);

// create ball
const material = new THREE.MeshLambertMaterial({color: 0x80FC66});
material.color.setRGB(material.color.r * 0.4, material.color.g * 0.4, material.color.b * 0.4);
const ball = new THREE.Mesh(new THREE.SphereGeometry(150, 64, 32), material)

// scene 
const scene = new THREE.Scene();
scene.fog = new THREE.Fog(0x808080, 2000, 4000);
scene.add(ambientLight);
scene.add(light);;
scene.add(ball);

const animate = function () {
    controls.update();
    requestAnimationFrame( animate );
    renderer.render( scene, camera );
};

animate();