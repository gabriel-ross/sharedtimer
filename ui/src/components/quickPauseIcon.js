import React from "react";
import {TiMediaPause} from 'react-icons/ti';

function QuickPauseIcon({timerId}) {
    function handleIconClick(timerId) {
    }
    return (
        <TiMediaPause className="clickable" onClick={() => handleIconClick(timerId)} />
    )
}

export default QuickPauseIcon