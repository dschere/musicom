
// rectangle overlay on the Tab beam that associates
// key/clicks to a given note/reset.
class TabCursor {
     x: number;
     y: number;
     
     constructor() {
         this.x = this.y = 0;
     }
}


const TabCursorId = "tabCursor";

class TabeditorView {
    vextab: any;
    elementId: string;
    
    vf : any;
    VexTab: any;
    Vex: any;

    constructor(vextab: any, elementId: string) {
        this.vextab = vextab;
        this.elementId = elementId;
    }

    positionTabSelector(x: number, y: number) {
        let elm = document.getElementById(TabCursorId);
        if (elm) {
            elm.style.display = "inline";
            elm.style.left = Number(Math.floor(x) - 8)+'px';
            elm.style.top = Number(Math.floor(y) - 8)+'px';
        }
    }
    
    hideTabSelector() {
        let elm = document.getElementById(TabCursorId);
        if (elm) {
            elm.style.display = "none";
        }
    }
    
    // use vextab to draw music score and tabature
    render(data: string){
        let vf = this.vextab.Vex.Flow;
            
        const renderer = new vf.Renderer(this.elementId, vf.Renderer.Backends.SVG);

        // Initialize VexTab artist and parser.
        const artist = new this.vextab.Artist(10, 10, 750, { scale: 0.8 });
        const tab = new this.vextab.VexTab(artist);

        tab.parse(data);
        artist.render(renderer);
    }
}

export default TabeditorView;
