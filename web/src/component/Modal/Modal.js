import React, { useState } from 'react';
import CloseIcon from '@material-ui/icons/Close';
import './Modal.scss';

const Modal = ({
  className = '',
  showState = true,
  isCloseButtonRequired = true,
  children,
  onClose,
  hideHandler,
  closeActionClassName = '',
  slideDirection = 'top',
  backDropClose = true,
  childData = null,
  mobileFullScreen = true,
  contentContainerClassName = ''
}) => {
  const [show, setShowState] = useState(showState);
  const [isModelOpen, setIsModelOpen] = useState(showState);

  const modalStateHandler = (status) => {
    setIsModelOpen(status);
    setShowState(status);
  };

  React.useEffect(() => {
    setShowState(showState);
    if (showState) {
      try {
        document.body.style.overflow = 'hidden';
      } catch (e) {
        // e
      }
    } else {
      try {
        document.body.style.overflow = 'auto';
      } catch (e) {
        // e
      }
    }
    setTimeout(() => {
      setIsModelOpen(showState);
    }, 200);
  }, [showState]);

  const closeModal = React.useCallback(() => {
    if (typeof onClose === 'function') {
      onClose();
    }
    if (typeof hideHandler === 'function') {
      hideHandler();
    }
    modalStateHandler(false);
  }, [onClose, hideHandler]);

  return show ? (
    <>
      {isModelOpen && (
        <div
          className='modal-backdrop'
          onClick={() => {
            if (backDropClose === false) return;
            closeModal();
          }}
        />
      )}
      <div
        className={`modal ${
          mobileFullScreen ? 'mobile-full-screen' : ''
        } ${slideDirection} ${isModelOpen ? 'open' : ''}`}
      >
        <div className={`${className} modal__dialog`}>
          <div className='modal__content'>
            {isCloseButtonRequired && (
              <button
                type='button'
                className={`close-btn ${closeActionClassName}`}
                onClick={closeModal}
              >
                <CloseIcon />
              </button>
            )}
            <div className={`content-container ${contentContainerClassName}`}>
              {typeof children === 'function' ? children(childData) : children}
            </div>
          </div>
        </div>
      </div>
    </>
  ) : (
    ''
  );
};

export default Modal;
