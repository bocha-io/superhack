using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using TMPro;


public class BattleInfo : MonoBehaviour
{
    [SerializeField] TextMeshProUGUI _text;

    [SerializeField] BottomPanelController _panel;

    public void Start(){
        StartCoroutine(ShowText("Trainer wants to battle"));
    }
    public IEnumerator ShowText(string text){
        _text.text = text;
        yield return new WaitForSeconds(3);
        // callback()
        _panel.ChangeState(PanelState.PickingAction);
    }

    public void ShowPermanentText(string text){
        _text.text = text;
    }

}
