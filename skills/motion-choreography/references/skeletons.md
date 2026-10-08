# Skeletons

Tested together in one page (vanilla JS, GSAP 3.15). Read only the sections the page uses. Adapt names, keep the structure.

Data shape used throughout:

```js
const W = 1180, H = 380, NW = 190, NH = 62;          // world size and node size, in px
const nodes = [{id:'gw', name:'api-gateway', kind:'envoy', icon:'shield-check', x:260, y:159}, ...];
const edges = [['web','gw'], ['gw','co'], ...];       // index = edge id
const steps = [{t:'Gateway roteia', d:'...', focus:['gw','co'], edges:[1], at:14}, ...];
// optional per step: bad:true (the problem step, drawn in --error), hold:1.6 (seconds to linger), back:[1,0] (edges to travel in reverse)
const N = Object.fromEntries(nodes.map(n => [n.id, n]));
```

## wires

Nodes are absolutely positioned inside `#world` at their data coordinates, so paths come straight from the data. No DOM measuring.

```js
svg.innerHTML = edges.map(([a,b], i) => {
  const A = N[a], B = N[b], x1 = A.x+NW, y1 = A.y+NH/2, x2 = B.x, y2 = B.y+NH/2, c = (x2-x1)/2;
  const d = `M${x1},${y1} C${x1+c},${y1} ${x2-c},${y2} ${x2},${y2}`;
  return `<path d="${d}" fill="none" stroke="var(--border)" stroke-width="1.5"/>
    <path id="w${i}" d="${d}" fill="none" stroke="var(--accent)" stroke-width="2" stroke-linecap="round"/>
    <circle id="p${i}" r="4.5" fill="var(--accent)" style="filter:drop-shadow(0 0 8px var(--accent))" opacity="0"/>`;
}).join('');
```

Each node holds a `.ring` child (absolute, `inset:-1px`, accent border plus soft glow `box-shadow`, `opacity:0`). Animating the ring's opacity is the "node lights up" effect.

## camera

```html
<div id="cam" class="relative overflow-hidden rounded-2xl border border-border h-[clamp(320px,48vw,500px)]">
  <div id="world"><svg id="wires" class="absolute inset-0 overflow-visible pointer-events-none"></svg><!-- nodes --></div>
  <div class="vignette absolute inset-0 pointer-events-none"></div>
  <div id="caption" class="absolute left-4 bottom-4 max-w-md pointer-events-none">...</div>
</div>
<style>
#world{position:absolute;left:0;top:0;transform-origin:0 0}
#world::before{content:"";position:absolute;inset:-2000px;background-image:radial-gradient(var(--grid) 1px,transparent 1px);background-size:24px 24px} /* grid moves with the camera */
.vignette{background:radial-gradient(120% 90% at 50% 45%,transparent 55%,var(--bg) 100%)}
</style>
```

```js
function fitAll(){ const cw = cam.clientWidth, ch = cam.clientHeight, s = Math.min(cw/W, ch/H)*0.94;
  return {s, x:(cw-W*s)/2, y:(ch-H*s)/2}; }
function shot(ids){                       // camera target for a set of node ids; null = fit all
  const all = fitAll(); if(!ids || RM) return {x:all.x, y:all.y, scale:all.s};
  const P = 60, b = ids.map(id => N[id]);
  const x0 = Math.min(...b.map(n=>n.x))-P, x1 = Math.max(...b.map(n=>n.x+NW))+P;
  const y0 = Math.min(...b.map(n=>n.y))-P, y1 = Math.max(...b.map(n=>n.y+NH))+P;
  const cw = cam.clientWidth, ch = cam.clientHeight - caption.offsetHeight - 16;  // frame above the caption
  const s = gsap.utils.clamp(all.s, Math.max(all.s*1.7, 0.85), Math.min(cw/(x1-x0), ch/(y1-y0)));
  return {scale:s, x:cw/2-(x0+x1)/2*s, y:ch/2-(y0+y1)/2*s};
}
// in the step loop (see timeline):
tl.to(world, {...shot(st.focus), duration:1.2, ease:'power3.inOut'}, L);
// depth of field
if(st.focus){ tl.to(outFocus, {filter:'blur(3px)', opacity:0.3, duration:0.9}, L)
                .to(inFocus,  {filter:'blur(0px)', opacity:1,   duration:0.9}, L); }
else tl.to(nodeEls, {filter:'blur(0px)', opacity:1, duration:0.9}, L);   // pull-back: everything sharp
```

The `Math.max(all.s*1.7, 0.85)` cap lets phones zoom to near real size while desktops stay subtle.

## trace

Inside the step loop, after the camera starts (offset about 0.6s so the move leads):

```js
st.edges.forEach((ei, k) => {
  const at = L+'+='+(0.6+k*0.6);
  tl.to('#w'+ei, {drawSVG:'100%', duration:0.7, ease:'power2.inOut'}, at)
    .to('#p'+ei, {opacity:1, duration:0.12}, at)
    .to('#p'+ei, {motionPath:{path:'#w'+ei, align:'#w'+ei, alignOrigin:[0.5,0.5]}, duration:0.7, ease:'power2.inOut'}, at)
    .to('#p'+ei, {opacity:0, duration:0.2}, L+'+='+(1.2+k*0.6))
    .to('#n-'+edges[ei][1]+' .ring', {opacity:1, duration:0.4}, L+'+='+(1.15+k*0.6));
});
// reverse travel (response going back): motionPath {..., start:1, end:0}
// linked view on the same timeline, e.g. waterfall bars:
spans.forEach((s,j) => { if(s.step === i) tl.to(bars[j], {scaleX:1, duration:Math.max(0.4, s.d/total*1.8), ease:'expo.out'}, L+'+=0.9'); });
```

Before building: `gsap.set('#wires path[id]', {drawSVG:'0%'})`, `gsap.set(bars, {scaleX:0})`, bars with `transform-origin:left`.

## timeline

```js
let tl;
function build(){
  const p = tl ? tl.progress() : 0, playing = tl && tl.isActive();
  tl && tl.kill();
  /* gsap.set(...) every animated thing to its rest state, world to fitAll() */
  tl = gsap.timeline({paused:true, defaults:{ease:'power3.out'}, onUpdate:sync});
  steps.forEach((st, i) => {
    const L = 's'+i; tl.addLabel(L);
    /* camera, depth of field, trace, linked views */
    tl.to({}, {duration: st.hold || 0.7});   // breathing room; longer hold on the "one thing"
  });
  tl.progress(p); sync(); if(playing) tl.play();
}
let rz; new ResizeObserver(() => { clearTimeout(rz); rz = setTimeout(build, 100); }).observe(cam);
```

## player

Markup (classes from the stack's component layer):

```html
<div class="flex flex-wrap items-center gap-x-4 gap-y-3">
  <div class="inline-flex items-center gap-0.5 p-1 rounded-full border border-border bg-muted/50">
    <button class="ib" id="prev" aria-label="Previous step"><i data-lucide="skip-back"></i></button>
    <button class="ib ib-main" id="play" aria-label="Play"><i data-lucide="play"></i></button>
    <button class="ib" id="next" aria-label="Next step"><i data-lucide="skip-forward"></i></button>
  </div>
  <div id="segs" class="basis-full sm:basis-0 flex-1 min-w-0 flex gap-1 py-2 cursor-pointer touch-none" role="slider" tabindex="0"></div>
</div>
```

```js
segs.innerHTML = steps.map(st => `<div class="seg"><i${st.bad?' style="background:var(--error)"':''}></i></div>`).join('');
const LB = i => i < steps.length ? tl.labels['s'+i] : tl.duration();
const current = () => { let c = 0; steps.forEach((_,i) => { if(tl.time() >= LB(i)-0.001) c = i; }); return c; };
let lastStep = -1, lastIcon;
function sync(){
  const c = current(), t = tl.time();
  fills.forEach((el,i) => el.style.transform = `scaleX(${gsap.utils.clamp(0,1,(t-LB(i))/(LB(i+1)-LB(i)))})`);
  if(c !== lastStep){ lastStep = c; /* update caption text, then: */
    gsap.fromTo('#caption > *', {y:8, opacity:0, filter:'blur(4px)'}, {y:0, opacity:1, filter:'blur(0px)', duration:0.5, stagger:0.05, ease:'expo.out', overwrite:true}); }
  const ic = tl.isActive() ? 'pause' : (tl.progress() === 1 ? 'rotate-ccw' : 'play');
  if(ic !== lastIcon){ lastIcon = ic; playBtn.innerHTML = `<i data-lucide="${ic}"></i>`; lucide.createIcons(); }
}
playBtn.onclick = () => { tl.isActive() ? tl.pause() : (tl.progress() === 1 ? tl.restart() : tl.play()); setTimeout(sync); };
next.onclick = () => tl.tweenTo(LB(current()+1));
prev.onclick = () => { const c = current(); tl.pause().seek(LB(tl.time()-LB(c) < 0.4 ? Math.max(0,c-1) : c)); sync(); };
function seekX(e){ const r = segs.getBoundingClientRect(), w = r.width/steps.length,
  x = gsap.utils.clamp(0, r.width-0.01, e.clientX-r.left), i = Math.floor(x/w);
  tl.pause().time(LB(i) + (x-i*w)/w*(LB(i+1)-LB(i))); sync(); }
segs.onpointerdown = e => { segs.setPointerCapture(e.pointerId); seekX(e); segs.onpointermove = seekX; };
segs.onpointerup = () => segs.onpointermove = null;
addEventListener('keydown', e => { if(e.key === ' '){ e.preventDefault(); playBtn.click(); }
  else if(e.key === 'ArrowRight') next.click(); else if(e.key === 'ArrowLeft') prev.click(); });
```

The play button's filled icon comes from `.ib-main svg{fill:currentColor;stroke-width:0}`. For `rotate-ccw`, reset to `fill:none;stroke-width:2.2`.

## derived text

Counters and status text come from the playhead, so they are right in any direction:

```js
// per entity, the state at the end of each step, e.g. 'RRRTX' (one char per step)
const stateAt = (e, c, dt, delay = 1.2) => (dt < delay && c > 0) ? e.st[c-1] : e.st[c];  // show the old state until the step's animation lands
// in sync(): const c = current(), dt = tl.time() - LB(c);
// readyEl.textContent = entities.filter(e => stateAt(e, c, dt) === 'R').length;
```

## reveal

```js
document.fonts.ready.then(() => {
  build();
  if(RM){ tl.progress(1); sync(); return; }
  const split = SplitText.create('h1', {type:'words', mask:'words'});
  gsap.from(split.words, {yPercent:110, opacity:0, filter:'blur(6px)', duration:0.9, stagger:0.045, ease:'expo.out'});
  gsap.from(['header p', '#cam', '#segs'], {y:14, opacity:0, filter:'blur(6px)', duration:0.8, stagger:0.07, ease:'expo.out', delay:0.15});
  gsap.delayedCall(1.2, () => tl.play());
});
```
