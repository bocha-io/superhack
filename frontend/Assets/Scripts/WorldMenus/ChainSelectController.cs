using System.Collections;
using System.Collections.Generic;
using Unity.VisualScripting;
using UnityEngine;
using TMPro;

public class ChainSelectController : OptionsController
{
    [SerializeField] TextMeshProUGUI _chainName;
    [SerializeField] WorldCanvasController _worldUI;

    public override void Start(){
        _chainName.text = Connection.Instance.chain;
        // switch (Connection.Instance.chain)
        // {
        //     case "zora":
        //     {
        //         options[0].gameObject.SetActive(false);
        //         options.RemoveAt(0);
        //         break;
        //     }
        //     case "base":
        //     {
        //         options[1].gameObject.SetActive(false);
        //         options.RemoveAt(1);
        //         break;
        //     }
        //     case "testnet":
        //     {
        //         options[2].gameObject.SetActive(false);
        //         options.RemoveAt(2);
        //         break;
        //     }
        //     case "optimism":
        //     {
        //         options[3].gameObject.SetActive(false);
        //         options.RemoveAt(3);
        //         break;
        //     }

        //     default:
        //         break;
        // }

        base.Start();
    }

    public override void Update(){
        previouslySelected = selected;
        // if(Input.GetKeyDown(KeyCode.D)){
        //     if (selected > 0){
        //         selected-=horizontalAugment;
        //     }                
        // }
        // if(Input.GetKeyDown(KeyCode.A)){
        //     if (selected < options.Count - horizontalAugment){
        //         selected+=horizontalAugment;
        //     }                
        // }
        // if(Input.GetKeyDown(KeyCode.S)){
        //     if (selected > 0){
        //         selected-=verticalAugment;
        //     }
        // }
        // if(Input.GetKeyDown(KeyCode.W)){
        //     if (selected < options.Count - verticalAugment){
        //         selected+=verticalAugment;
        //     }

        // }
        
        // if(Input.GetKeyDown(KeyCode.F)){
        //     options[selected].Execute();
        // }
        if(Input.GetKeyDown(KeyCode.G)){
            if (canGoBack)
                GoBack();
        }
        if (previouslySelected != selected)
            Select();
    }

    public  override void GoBack(){
        _worldUI.Close();
    }

}
