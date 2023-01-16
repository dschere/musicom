import {vextab, RenderQ_EventShim} from "./view/myvextab.js"

const Renderer = vextab.Vex.Flow.Renderer;

export class TrackEditor {
    evthandler: any
    renderer: any

    constructor(elementId: string){
       //!!!!!!! This alters vetflow.Factory.renderQ 
       // see myvextab.js for details.
       this.evthandler = new RenderQ_EventShim();
       this.renderer = 
           new Renderer(elementId, Renderer.CANVAS );
        
    }

    public render() {
         
    }
}


