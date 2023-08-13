using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class OptionsController : MonoBehaviour
{
    public List<MenuOption> options;
    public int selected, previouslySelected;
    public bool canGoBack;
    [SerializeField] BottomPanelController _parentPanel;
    
    public bool horizontal = true;
    public int horizontalAugment = 1;
    public int verticalAugment = 2;

    public bool enable = true;

    public virtual void Start(){
       if (!horizontal) {
            horizontalAugment = 2;
            verticalAugment = 1;
       }
    }

    public virtual void Setup(List<string> names){
        
    }

    public void OnEnable(){
        Select();
    }

    public virtual void GoBack(){
        if(canGoBack){
            _parentPanel.ChangeState(PanelState.PickingAction);
        }
    }

    public virtual void Update(){
        if (!enable) return;

        previouslySelected = selected;
        if(Input.GetKeyDown(KeyCode.D)){
            if (selected < options.Count - horizontalAugment){
                selected+=horizontalAugment;
            }                
        }
        if(Input.GetKeyDown(KeyCode.A)){
            if (selected > 0){
                selected-=horizontalAugment;
            }                
        }
        if(Input.GetKeyDown(KeyCode.S)){
            if (selected < options.Count - verticalAugment){
                selected+=verticalAugment;
            }                
        }
        if(Input.GetKeyDown(KeyCode.W)){
            if (selected > 0){
                selected-=verticalAugment;
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
