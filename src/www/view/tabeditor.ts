import {RenderQ_EventShim, VexTab, Vex} from "./myvextab.js";


const VF = VexTab.Vex.Flow;

export class TabatureEditor {
    vexEvents : any

    private vexFlowFactoryEventHandler(e: any){
        // capture events from vextab
        console.log(e.details.className);
        console.log(e.details.obj);
    }

    constructor() {
        this.vexEvents = new RenderQ_EventShim();
        this.vexEvents.addEventListener(
            'vexFlowFactoryEvent',
            this.vexFlowFactoryEventHandler 
        );
    }

    render(data: string, elementId: string){

        const renderer = new VF.Renderer(elementId,
            VF.Renderer.Backends.CANVAS);

        // Initialize VexTab artist and parser.
        const artist = new VexTab.Artist(10, 10, 750, { scale: 0.8 });
        const tab = new VexTab.VexTab(artist);

        tab.parse(data);
        artist.render(renderer);
    }

}

