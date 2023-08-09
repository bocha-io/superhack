using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class OptionsController : MonoBehaviour
{
    public List<MenuOption> options;
    public int selected, previouslySelected;
    public bool canGoBack;
    [SerializeField] BottomPanelController _parentPanel;
    
    public void Setup(List<string> names){
        
    }

    public void OnEnable(){
        Select();
    }

    public void GoBack(){
        if(canGoBack){
            _parentPanel.ChangeState(PanelState.PickingAction);
        }
    }

    public void Update(){
        previouslySelected = selected;
        if(Input.GetKeyDown(KeyCode.D)){
            if (selected < options.Count - 1){
                selected++;
            }                
        }
        if(Input.GetKeyDown(KeyCode.A)){
            if (selected > 0){
                selected--;
            }                
        }
        if(Input.GetKeyDown(KeyCode.S)){
            if (selected < options.Count - 2){
                selected+=2;
            }                
        }
        if(Input.GetKeyDown(KeyCode.W)){
            if (selected > 1){
                selected-=2;
            }                
        }
        
        if(Input.GetKeyDown(KeyCode.F)){
            options[selected].Execute();
        }
        if(Input.GetKeyDown(KeyCode.G)){
            if (canGoBack)
                GoBack();
        }
        if (previouslySelected != selected)
            Select();
    }

    public virtual void Select(){
        for (int i=0; i< options.Count; i++){
            options[i].Select(i==selected);
        }
    }


}
