using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using DG.Tweening;
public class BochamonOption : MenuOption
{
    public Bochamon _bochamon;
    public Image _bochamonImage;
    [SerializeField] BochamonInfo _bochaInfo;
    // [SerializeField] BochamonInfo _bochamonInfo;
    public override void Select(bool s)
    {
        _selection.gameObject.SetActive(s);
        if (s) {
            _bochaInfo.Setup(_bochamon);
            _selection.transform.localScale = new Vector3 (0.8f, 0.9f, 1);
            _selection.transform.DOScale(new Vector3(1,1,1), 0.1f);
        }

        // base.Select(s);
        // _attackInfo.Setup(_move);
    }

    public void Setup(Bochamon bochamon){
        _bochamon = bochamon;
        Debug.Log("Bochamon has " + bochamon.currentHp.ToString());
        if (bochamon.currentHp > 0)
            _bochamonImage.sprite = bochamon.sprite;
        else _bochamonImage.sprite = bochamon.greySprite;
        base.Setup(bochamon.bochaName);
    }
}
