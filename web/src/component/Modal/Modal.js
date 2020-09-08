/* eslint-disable operator-linebreak */
import React, { useState } from 'react';
import './Modal.scss';

const CloseIcon =
  'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAE6UlEQVR4Xu2aV88dNRBAT+hNPPB/QhUgRA0CQu+htwgEAhIQoSMQvfeS0JtAIET9R4gHegk6sEaXq7u7Htv7RZ++69e1xzPH41l7POtY423dGrefJYClB6xxAsstsMYdYBkES7bAAcB24FDgKeDmXehFuwGPABcDXwJnAj9G9IkC2B/4HFg/M4kQrgJ2RiZu0FfjXwLOm5H1HXAs8FOu/AiA/YDPgEMWCH8WuHwFIWj8K8A5C3T5toPwcw6EXAAa/ylw2IDQF4BLVgDC7sCrwFkDunwDHAeMQsgBsG9n/OEZRF/u9uNfGX1Lumj868AZGYO/Bo4fg5AD4H3g5IwJU5fXgAuA1hD26ILvaQFd3gNOHeqfA+B3wMkj7Q3gfODPyKCBvs7/JnBKUJ7B0MDd23IAGGld0WhTYYNULYQ9gbeCXph0fRq4ohZAieulOd/pgtUfUXpdf41/G9hQMN5YoRcObsUcD3Bug49ufXqBIu5Dg1YUwl6AAE8smNNf5EVjxis3F0CCMPb76dP1Q2AjYDzJaRovOKN4tLllN+UYHwVg/6EDyJiiHwNG8N9GOu4N+OfxRBdtzwOXRs4iEQ9IygjhxW5/RRX0MGUk/7VnoMbrLcdEBQPPdAEvdCQvAZA8QdoXFijqXcKgNg9hH+Aj4OgCmcX3kVIAaft4B3C/RdsXwEnAL91AT5tukSOjgoDHgWsKxv0zpAZAGi/9ywoU8PqaIvwnwBEFMrwKby4Y99+QWgAJwhNjB44eJT2vq0POPWNexEPADTXGt/CA2fkfA66uVShz/APATZl9B7u18IDZCR4Grmuh2ICMe4FbWs3RGoB6PQhc30rBOTl3AVtbyp4CgPrdD9zYUlFgG3B7Y5nVf4Ehfe5pmDDVcAE0b1N5QFL0TmBLpdaOv7tSRu/wqQE48R3AbYUG3AroSZO1qQF4bzBZWpJQ0Whvdub8Q+f7CK0pASjbS1Op8ckOZXjcngTCVACU68qXXJYWLeBkKfcpACjTm6IZmZZtEgitAUxlfAIZTniMrUBLAMp6rgtaY/PWfHcOb59NYkIrAMopzQ2UwGj2FtkCgDJMR/kuGG0+pdlK/hRFKbB5BWsB1BifUtfqVJpj9OHjyprtUAPAsSpgFjbaTK/7i0yPFove+nNlmpESQlErBeC40lRY3+OpENwS5xZYUgyhBIBjnuwKIqK6jr0c10BQJytVQi0KoMZ4n9YsZxl7Nq95fDE3GUrLRQFIefC1tQe/RVW69pjxabgQjBNnh5bz386hNHkEQKnxOzrjo8/kQnDLDJXC9PExQXttDrxcAKXG19YI+CotBMvfou3RnARtDoD7ClPQFjXowtGVnzc0Uhc0P3Y0iZoD4AfgwCB+ixp03Vrj07SlEL4HDhrSPQfAB8EKDYsadNlWxs9CiBZpqIt1Cb0tB0BOjWCa4N3O+Gg1SK6D6Qn+UQaN6oR9BZzQokxOeUNVokn50lKYXONnPcE/y1C5XFaNoAJzPCBNbLmZpbIHL9DYig7rh6Za+fkpLdwSwqIawOwq0SgA+y8qljZG6JIrZXyCsah2cNJi6TSx5fK6+1HdfvRWl1v8FHX3sf5CMFdoPaKVJ26L7ErxEg+YVcgavl1l+DwYT425x+z/jY3EgLHVWJXflwBW5bI1VHrpAQ1hrkpRSw9YlcvWUOm/Aa812kErBGjPAAAAAElFTkSuQmCC';

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
                <img src={CloseIcon} alt='close' />
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
