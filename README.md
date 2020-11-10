# pixelArtVectorizer
Vectorize Pixel Arts in Go
Convert pixel art to SVG vector image

Basead on my previous project SVG Vectorize in .net framework https://github.com/michaelrbk/pixelart-vectorizer

The old version was a proof of concept and the creation of the core algorith. This new version in go langhas the goal to create a console app to convert a pixel art to a new image file vectorized. With several improvements aiming on make the algorith usable in other aplications, maybe a emulator.


* Source file
<p align="center">
<img src="./results/0.source.svg" height="200">
</p>
* Generating a graph with similar color
<p align="center">
<img src="./results/1.genGraph.svg" height="200">
</p>
* Solving ambiquities in the graph's connections 
<p align="center">
<img src="./results/2.solveAmbiguities.svg" height="200">
</p>

Copyright [2020] [Michael de Rezende Becker]

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
