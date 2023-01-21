
import TabeditorView from './tabeditor_view';


class TabEditorController {
    te_view: any;
    vextab: any;
    elementId: string;
    movements: []; // notes/rests + meta data
    
    
    on_tabnoteRendered(evt: any) {
        console.log("on_tabnoteRendered");
        if (this.vextab.VexEvents.current_element === this.elementId) {
            let tabnoteQ = evt.detail.eventQueue;
            console.log(tabnoteQ);
            //let data = tabnoteQ.pop();
            //if (data) {
            //    this.te_view.positionTabSelector(data[0].x, data[0].y);
            //}
        }
    }

    constructor(vextab: any, elementId: string) {
         // This allows for multiple tab controllers (one for each track) 
         
         if ( document.getElementById(elementId) === null ) {
             throw "No dom element with id=" + elementId + " found";
         }
         
         this.te_view = new TabeditorView(vextab, elementId);
         
         this.vextab = vextab; 
         this.elementId = elementId;
         this.movements = [];

         // setup event handler for when tabiture is rendered we
         // can capture what and where it was rendered for interactive
         // rendering
         vextab.VexEvents.addEventListener('vexTabNoteRender', (evt: any)=> {
             this.on_tabnoteRendered(evt);
         }); 

    }
    
    updateView(data: string) {
         
         this.vextab.VexEvents.current_element = this.elementId;

         let elm = document.getElementById(this.elementId);

         if (elm) { 
             elm.innerHTML = ''; // remove any previous svg
         }
         setTimeout( ()=> {
             this.te_view.render(data);
         }, 0);
    }

}

export default TabEditorController;
