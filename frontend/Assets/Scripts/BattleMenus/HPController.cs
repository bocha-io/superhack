using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using TMPro;
using DG.Tweening;

public class HPController : MonoBehaviour
{
    [SerializeField] TextMeshProUGUI _maxHpText;
    [SerializeField] TextMeshProUGUI _currentHpText;
    [SerializeField] Image _hpbar;
    int _maxHp;
    int _currentHp;

    Camera _mainCamera;
    void Start(){
        _mainCamera = Camera.main;
    }

    public void Setup(int maxHp, int currentHp){
        _maxHp = maxHp;
        _currentHp = currentHp;
        _maxHpText.text = maxHp.ToString();
        SetHp(_currentHp);
    }

    public void SetHp(int currentHp){
        _currentHpText.text = currentHp.ToString();
        float size = (float)_currentHp/(float)_maxHp;
        _hpbar.gameObject.transform.DOScaleX(size, 1f);// = new Vector3(size, 1, 1);
        if (size > 0.5f){
            _hpbar.color = Color.green;
        } else if (size < 0.15f){
            _hpbar.color = Color.red;
        } else {
            _hpbar.color = Color.yellow;
        }
    }

    public void ApplyDamage(int damage){
        _currentHp-=damage;
        if (_currentHp < 0){
            _currentHp = 0;
        }
        SetHp(_currentHp);
    }

}

