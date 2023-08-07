using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class OptionsController : MonoBehaviour
{
    public List<MenuOption> options;
    public int selected;

    // public void Setup(List<MenuOption> load){

    // }

    public void Update(){
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

        for (int i=0; i< options.Count; i++){
            options[i].Select(i==selected);
        }
    }


}
