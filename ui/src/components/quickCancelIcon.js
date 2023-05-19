import React from "react";
import {TiMediaStop} from 'react-icons/ti';

function QuickViewCancelIcon({timerId}) {
    function handleIconClick(timerId) {
        
    }
    return (
        <TiMediaStop className="clickable" onClick={() => handleIconClick(timerId)} />
    )
}

export default QuickViewCancelIcon