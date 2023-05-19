import React from "react";
import {TiMediaPlay} from 'react-icons/ti';

function QuickStartIcon({timerId}) {
    function handleIconClick(timerId) {
        
    }
    return (
        <TiMediaPlay className="clickable" onClick={() => handleIconClick(timerId)} />
    )
}

export default QuickStartIcon