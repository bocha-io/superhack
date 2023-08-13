using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using TMPro;

public class ChainSelector : MonoBehaviour
{
    [SerializeField] Image _background;
    [SerializeField] Image _bochamon;
    [SerializeField] TextMeshProUGUI _text;

    [SerializeField] Sprite _zoraBackground;
    [SerializeField] Sprite _zoraBocha;

    [SerializeField] Sprite _baseBackground;
    [SerializeField] Sprite _baseBocha;

    [SerializeField] Sprite _optimismBackground;
    [SerializeField] Sprite _optimismBocha;

    [SerializeField] Sprite _testnetBackground;
    [SerializeField] Sprite _testnetBocha;
    // Start is called before the first frame update

    int currentlySelected = 0;

    public void Next(){
        currentlySelected++;
        if (currentlySelected > 3)
            currentlySelected = 0;
        Select(currentlySelected);
    }

    public void Previous(){
        currentlySelected--;
        if (currentlySelected < 0)
            currentlySelected = 3;
        Select(currentlySelected);
    }

    public void Select(int select){
        switch(select){
            case 0:
            // Zora
            {
                _text.text = "Zora";
                _background.sprite = _zoraBackground;
                _bochamon.sprite = _zoraBocha;
                break;
            }
            case 1:
            // Base
            {
                _text.text = "Base";
                _background.sprite = _baseBackground;
                _bochamon.sprite = _baseBocha;
                break;
            }
            case 2:
            // Optimism
            {
                _text.text = "Optimism";
                _background.sprite = _optimismBackground;
                _bochamon.sprite = _optimismBocha;
                break;
            }
            case 3:
            // Testnet
            {
                _text.text = "Testnet";
                _background.sprite = _testnetBackground;
                _bochamon.sprite = _testnetBocha;
                break;
            }   
            default:
                break;
        }
    }

}
