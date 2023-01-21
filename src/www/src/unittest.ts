import TabEditorController from './comp/tabeditor_controller';

function test_tab_editor(vextab: any, elementId: string) {
    console.log("test_tab_editor called");
    let tc = new TabEditorController( vextab, elementId );
    let tab = `
  tabstave notation=true key=A time=4/4

  notes :w ##
  notes :q =|: (5/2.5/3.7/4) :8 7-5h6/3 ^3^ 5h6-7/5 ^3^ :q 7V/4 |
  notes :8 t12p7/4 s5s3/4 :8 3s:16:5-7/5 :q p5/4 |
  notes :w ##
  
  
`
// text :w, |#segno, ,|, :hd, , #tr
    tc.updateView(tab);
}

export default test_tab_editor;
