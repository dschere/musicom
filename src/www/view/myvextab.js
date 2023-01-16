//import * as _vextab from "node_modules/vextab/releases/main.dev.js"

import vextab from "node_modules/vextab/releases/main.dev.js"


export const VexTab = vextab.VexTab;
export const Vex = vextab.Vex;

/*
    A work around for the lack of events provided
    by vexflow. I need to be able to trap exactly what
    and where any note is being drawn to provide the
    basis for an editor.
    
    The Factory.renderQ has the info I need. Unfortunately
    I need javascript introspection so I can't use
    a minified javascript. 
*/
export class RenderQ_EventShim extends EventTarget {
    
    constructor() {
        this.renderQ = [];
        vextab.Vex.Factory.renderQ = this;  
    }

    emitFactoryEvent(className, obj) {
        this.dispatchEvent(
           new CustomEvent('vexFlowFactoryEvent', { detail: {
             className,
             obj
           }})
        );
    }
   
    // when factory.draw calls this.renderQ.forEach
    // events will be fired.
    forEach( callback ) {
        var emitFactoryEvent = this.emitFactoryEvent;
        this.renderQ.forEach((obj) => {
            emitFactoryEvent(obj.constructor.className, obj)
        });
    }

    push(obj) {
        this.renderQ.push(obj);
    }
}
