import { useEffect, useRef, useState } from 'react';

function App() {
  const [isLoaded, setIsLoaded] = useState(false);
  const planeRef = useRef(null);
  const cloudRefs = useRef([]);
  const bannerRef = useRef(null);

  useEffect(() => {
    setIsLoaded(true);
  }, []);
  return (
    <div className='container'>
      {/* SVG Animation */}
      <svg viewBox='0 0 1920 1080' preserveAspectRatio='xMidYMid slice'>
        {/* Sun */}
        <defs>
          <radialGradient
            id='sunGradient'
            cx='50%'
            cy='50%'
            r='50%'
            fx='50%'
            fy='50%'
          >
            <stop offset='0%' stopColor='#ffef5e' />
            <stop offset='90%' stopColor='#f59e0b' />
          </radialGradient>
          <filter id='sunGlow' x='-50%' y='-50%' width='200%' height='200%'>
            <feGaussianBlur stdDeviation='15' result='blur' />
            <feComposite in='SourceGraphic' in2='blur' operator='over' />
          </filter>
        </defs>

        <circle
          cx='1700'
          cy='200'
          r='100'
          fill='url(#sunGradient)'
          filter='url(#sunGlow)'
          className='animate-pulse-sun'
        />

        {/* Distant Mountains */}
        <path
          d='M0,800 L200,650 L400,750 L600,600 L800,700 L1000,620 L1200,750 L1400,650 L1600,720 L1800,680 L1920,750 L1920,1080 L0,1080 Z'
          fill='#36454f'
          opacity='0.3'
        />

        {/* Clouds */}
        {[...Array(12)].map((_, i) => {
          const scale = 0.6 + Math.random() * 0.8;
          const yPos = 100 + Math.random() * 400;

          return (
            <g
              key={i}
              ref={(el) => (cloudRefs.current[i] = el)}
              className='cloud'
              style={{ "{{" }}
                transform: `translate(${
                  -300 + i * 200
                }px, ${yPos}px) scale(${scale})`,
                opacity: 0.8 + Math.random() * 0.2,
              }}
            >
              <path
                d='M25,60 
                   Q40,40 60,50 
                   Q80,5 110,40 
                   Q135,10 160,40 
                   Q190,15 220,50
                   Q245,25 260,60
                   Q280,40 300,70
                   Q300,100 280,110
                   Q285,130 260,140
                   Q260,155 240,160
                   Q220,180 190,160
                   Q170,180 140,160
                   Q120,190 90,160
                   Q65,175 40,150
                   Q20,160 10,140
                   Q0,120 10,100
                   Q0,80 15,70
                   Q10,50 25,60'
                fill='white'
                filter='drop-shadow(0px 10px 5px rgba(0, 0, 0, 0.1))'
              />
            </g>
          );
        })}

        {/* Birds */}
        {[...Array(8)].map((_, i) => {
          const scale = 0.5 + Math.random() * 0.5;
          const yPos = 200 + Math.random() * 300;
          const speed = 20 + Math.random() * 20;
          const delay = i * 2;

          return (
            <g
              key={`bird-${i}`}
              style={{ "{{" }}
                animation: `fly-bird ${speed}s linear ${delay}s infinite`,
                transform: `translate(${
                  -100 + i * 50
                }px, ${yPos}px) scale(${scale})`,
              }}
            >
              <path
                d='M0,0 Q10,-15 20,0 Q30,-15 40,0'
                stroke='#333'
                strokeWidth='2'
                fill='none'
                className='animate-flap-wings'
              />
            </g>
          );
        })}

        {/* Plane with Banner */}
        <g
          ref={planeRef}
          className={`plane-group ${isLoaded ? 'animate-fly-plane' : ''}`}
          style={{ "{{" }} transform: 'translate(-400px, 300px)' }}
        >
          {/* Plane */}
          <g className='plane'>
            {/* Plane Shadow */}
            <ellipse
              cx='70'
              cy='125'
              rx='60'
              ry='10'
              fill='rgba(0,0,0,0.1)'
              filter='blur(5px)'
            />

            {/* Plane Body */}
            <g>
              {/* Main Body */}
              <path
                d='M40,100 
                   L120,85 
                   C130,85 140,90 145,100 
                   C140,110 130,115 120,115 
                   L40,100 Z'
                fill='#3b82f6'
                stroke='#1d4ed8'
                strokeWidth='2'
              />

              {/* Cockpit */}
              <path
                d='M120,85 
                   C130,85 140,90 145,100 
                   C140,110 130,115 120,115 
                   C115,110 115,90 120,85 Z'
                fill='#dbeafe'
                stroke='#1d4ed8'
                strokeWidth='1.5'
              />

              {/* Windows */}
              <path
                d='M60,95 L70,93 L70,107 L60,105 Z 
                   M80,91 L90,89 L90,111 L80,109 Z'
                fill='#dbeafe'
                stroke='#1d4ed8'
                strokeWidth='1'
              />

              {/* Top Wing */}
              <path
                d='M70,85 
                   C70,85 90,55 110,55 
                   L110,85 
                   L70,85 Z'
                fill='#60a5fa'
                stroke='#1d4ed8'
                strokeWidth='1.5'
              />

              {/* Bottom Wing */}
              <path
                d='M70,115 
                   C70,115 90,145 110,145 
                   L110,115 
                   L70,115 Z'
                fill='#60a5fa'
                stroke='#1d4ed8'
                strokeWidth='1.5'
              />

              {/* Tail */}
              <path
                d='M40,100 
                   L20,80 
                   L30,100 
                   L20,120 
                   L40,100 Z'
                fill='#60a5fa'
                stroke='#1d4ed8'
                strokeWidth='1.5'
              />

              {/* Propeller Hub */}
              <circle
                cx='145'
                cy='100'
                r='5'
                fill='#94a3b8'
                stroke='#64748b'
                strokeWidth='1.5'
              />

              {/* Propeller Animation */}
              <g
                className='animate-spin-propeller'
                style={{ "{{" }} transformOrigin: '145px 100px' }}
              >
                <path
                  d='M145,100 L165,80 M145,100 L165,120'
                  stroke='#94a3b8'
                  strokeWidth='3'
                  strokeLinecap='round'
                />
              </g>
            </g>
          </g>

          {/* Banner Strings */}
          <path
            d='M30,110 C20,130 0,150 -20,160 
               M30,90 C20,70 0,50 -20,40'
            stroke='#f97316'
            strokeWidth='2'
            strokeDasharray='5,5'
            fill='none'
          />

          {/* Banner */}
          <g
            ref={bannerRef}
            className='animate-wave-banner'
            style={{ "{{" }} transformOrigin: '-20px 100px' }}
          >
            <rect
              x='-320'
              y='40'
              width='300'
              height='120'
              rx='10'
              fill='#f97316'
              filter='drop-shadow(0px 10px 8px rgba(0, 0, 0, 0.2))'
            />

            {/* Banner Folds */}
            <path
              d='M-320,60 C-280,70 -240,50 -200,70 C-160,90 -120,60 -80,70 C-40,80 -20,60 -20,60'
              stroke='rgba(255,255,255,0.3)'
              strokeWidth='2'
              fill='none'
            />
            <path
              d='M-320,100 C-280,110 -240,90 -200,110 C-160,130 -120,100 -80,110 C-40,120 -20,100 -20,100'
              stroke='rgba(255,255,255,0.3)'
              strokeWidth='2'
              fill='none'
            />

            {/* Banner Text */}
            <text
              x='-170'
              y='110'
              fontFamily='Arial'
              fontSize='50'
              fontWeight='bold'
              fill='white'
              textAnchor='middle'
              filter='drop-shadow(0px 2px 2px rgba(0, 0, 0, 0.3))'
            >
              React Go!
            </text>
          </g>
        </g>
      </svg>

      <style jsx global>{`
        @keyframes fly-plane {
          0% {
            transform: translate(-400px, 300px);
          }
          100% {
            transform: translate(2200px, 300px);
          }
        }

        @keyframes spin-propeller {
          0% {
            transform: rotate(0deg);
          }
          100% {
            transform: rotate(360deg);
          }
        }

        @keyframes wave-banner {
          0%,
          100% {
            transform: skewY(0deg);
          }
          25% {
            transform: skewY(3deg);
          }
          75% {
            transform: skewY(-2deg);
          }
        }

        @keyframes fly-bird {
          0% {
            transform: translate(-100px, var(--bird-y, 300px))
              scale(var(--bird-scale, 0.5));
          }
          100% {
            transform: translate(2000px, var(--bird-y, 300px))
              scale(var(--bird-scale, 0.5));
          }
        }

        @keyframes flap-wings {
          0%,
          100% {
            d: path('M0,0 Q10,-15 20,0 Q30,-15 40,0');
          }
          50% {
            d: path('M0,0 Q10,5 20,0 Q30,5 40,0');
          }
        }

        .animate-fly-plane {
          animation: fly-plane 20s linear infinite;
        }

        .animate-spin-propeller {
          animation: spin-propeller 0.1s linear infinite;
        }

        .animate-wave-banner {
          animation: wave-banner 4s ease-in-out infinite;
        }

        .animate-flap-wings {
          animation: flap-wings 0.5s ease-in-out infinite;
        }

        body {
          margin: 0;
          padding: 0;
          overflow: hidden;
          background: #87ceeb;
        }
      `}</style>
    </div>
  );
}

export default App;
