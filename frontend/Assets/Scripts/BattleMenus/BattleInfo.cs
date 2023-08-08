using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using TMPro;


public class BattleInfo : MonoBehaviour
{
    [SerializeField] TextMeshProUGUI _text;

    [SerializeField] BottomPanelController _panel;

    public void OnEnable(){
        StartCoroutine(ShowText());
    }
    public IEnumerator ShowText(){
        yield return new WaitForSeconds(1);
        // callback;'
        _panel.ChangeState(PanelState.PickingAction);
    }

}
