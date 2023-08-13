using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using TMPro;

public class TimerController : MonoBehaviour
{
    [SerializeField] TextMeshProUGUI _text;
    float currentTime=0;
    bool enable = false;
    public void Reset(bool e){
        currentTime = 0f;
        enable = e;
        if (!e){
            _text.text = "";
        }
    }

    void Update(){
        if (!enable) return;

        currentTime += Time.deltaTime;
        int seconds  = System.Convert.ToInt32( currentTime % 60);
        _text.text = (60 - seconds).ToString();

        if (currentTime >= 60){
            _text.text = "0";
            enable = false;
        }
    }





}
