/*
Snake oil code for emitting events associated with the rendering of the sheet music 
*/


function computeBeats(duration, dotted, isTriplet) {
    if (duration === 'w') { return 4; }
    
    if (isTriplet) {
        if (duration === 'h') { return 1.333333; }
        if (duration === 'q') { return 0.666666; }
        let n = Math.pow(0.5, parseInt(duration)/8);
        return (n*2)/3.0;
    } else {
        if (duration === 'h') { return 2 + ((dotted) ? 1  : 0); }
        if (duration === 'q') { return 1 + ((dotted) ? 0.5: 0); }
        let n = Math.pow(0.5, parseInt(duration)/8);
        return n + ((dotted) ? (n*0.5): 0);
    }
}



VexEvents = new EventTarget();
VexEvents.current_element = "";
VexEvents.tab_note_queue   = []; // only notes + fingering poistion 
VexEvents.stave_note_queue = []; // both rests/notes but no finger positions



vextab.Vex.Flow.StaveNote.prototype.base_draw = vextab.Vex.Flow.StaveNote.prototype.draw;
function MyStaveNoteDraw() {
    this.base_draw();

    let isRest   = this.noteType === "r";
    //TODO see if we can get rid of stupid magic number and find out why its
    // off by 10
    let tab_x    = (this.getCenterGlyphX() * this.getStave().getContext().state.scale.x) - 10.0;
    let isDotted = this.dots > 0;
    // is this note part of a triplet?
    let isTriplet = this.tickMultiplier.denominator === 3;
    
    let beats = computeBeats(this.duration, this.isDotted(), isTriplet);
     
    VexEvents.stave_note_queue.push({
        isRest,
        tab_x,
        beats,
        "StaveNote": this,
        fingerPos: null
    });
}
vextab.Vex.Flow.StaveNote.prototype.draw = MyStaveNoteDraw;

vextab.Vex.Flow.TabNote.prototype.base_drawPositions = vextab.Vex.Flow.TabNote.prototype.drawPositions;


function myDrawPositions() {
    // call original drawPositions to do the work.    
    this.base_drawPositions();

    //const ys = this.ys;
    const fingerPos = this.getPositions();
    
    VexEvents.tab_note_queue.push({
         fingerPos
    });    
}
vextab.Vex.Flow.TabNote.prototype.drawPositions = myDrawPositions;



vextab.Artist.prototype.base_render = vextab.Artist.prototype.render;
function myArtistRender(renderer) {
    this.base_render(renderer); // calls myDrawPositions 
    // after page render
    setTimeout(()=> { 
        var notes = [];        
        var stave_note_idx = 0;
        
        for( ; stave_note_idx < VexEvents.stave_note_queue.length; stave_note_idx++ ) {
            let note = VexEvents.stave_note_queue[stave_note_idx];
            if (!note.isRest) {
                let tn = VexEvents.tab_note_queue.pop();
                note.fingerPos = tn.fingerPos;
            }
            notes.push(note);
        }  
        
	var evt = { detail : {
	     'eventQueue': notes
	}};
	VexEvents.dispatchEvent( new CustomEvent('vexTabNoteRender',evt));
        
    }, 0);
}
vextab.Artist.prototype.render = myArtistRender;



vextab.VexEvents = VexEvents;


/*

let VexFlowFactoryDraw = vextab.Vex.Flow.Factory.prototype.draw;

vextab.Vex.Flow.Factory.prototype.draw = function() {
    console.log('cutom draw called!');
    for(let i =0; i < this.renderQ.length; i ++) { 
        let obj = this.renderQ[i];
        let className = obj.constructor.className;
        let evt = new CustomEvent(
            'vexFlowFactoryEvent', 
            { detail: {
                 className,
                 obj
                }
            }
        );
        VexEvents.dispatchEvent(evt);
    }
    VexFlowFactoryDraw(); // call original draw   
};

*/
