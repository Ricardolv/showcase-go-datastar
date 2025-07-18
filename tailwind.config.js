/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./internal/templates/**/*.templ",
    "./web/static/js/**/*.js",
    "./cmd/**/*.go",
    "./internal/**/*.go"
  ],
  theme: {
    extend: {
      colors: {
        // Paleta CEAR: Branco, Preto, Laranjado
        primary: {
          50: '#fff7ed',    // Laranja muito claro
          100: '#ffedd5',   // Laranja claro
          200: '#fed7aa',   // Laranja suave
          300: '#fdba74',   // Laranja médio claro
          400: '#fb923c',   // Laranja médio
          500: '#f97316',   // Laranja principal CEAR
          600: '#ea580c',   // Laranja escuro
          700: '#c2410c',   // Laranja muito escuro
          800: '#9a3412',   // Laranja profundo
          900: '#7c2d12',   // Laranja máximo
          950: '#431407'    // Laranja quase preto
        },
        
        secondary: {
          50: '#f9fafb',    // Branco suave
          100: '#f3f4f6',   // Cinza muito claro
          200: '#e5e7eb',   // Cinza claro
          300: '#d1d5db',   // Cinza médio claro
          400: '#9ca3af',   // Cinza médio
          500: '#6b7280',   // Cinza
          600: '#4b5563',   // Cinza escuro
          700: '#374151',   // Cinza muito escuro
          800: '#1f2937',   // Preto suave
          900: '#111827',   // Preto
          950: '#030712'    // Preto profundo
        }
      },
      
      fontFamily: {
        sans: ['Inter', 'system-ui', '-apple-system', 'sans-serif'],
        mono: ['JetBrains Mono', 'Monaco', 'Consolas', 'monospace']
      },
      
      boxShadow: {
        'orange': '0 1px 3px 0 rgba(249, 115, 22, 0.1), 0 1px 2px 0 rgba(249, 115, 22, 0.06)',
        'orange-md': '0 4px 6px -1px rgba(249, 115, 22, 0.1), 0 2px 4px -1px rgba(249, 115, 22, 0.06)',
        'orange-lg': '0 10px 15px -3px rgba(249, 115, 22, 0.1), 0 4px 6px -2px rgba(249, 115, 22, 0.05)',
        'orange-xl': '0 20px 25px -5px rgba(249, 115, 22, 0.1), 0 10px 10px -5px rgba(249, 115, 22, 0.04)',
        'orange-2xl': '0 25px 50px -12px rgba(249, 115, 22, 0.25)',
        'orange-inner': 'inset 0 2px 4px 0 rgba(249, 115, 22, 0.06)'
      },
      
      animation: {
        'bounce-in': 'bounceIn 0.6s ease-out',
        'slide-up': 'slideUp 0.4s ease-out',
        'fade-in': 'fadeIn 0.3s ease-out',
        'pulse-orange': 'pulseOrange 2s cubic-bezier(0.4, 0, 0.6, 1) infinite',
        'spin-slow': 'spin 3s linear infinite'
      },
      
      keyframes: {
        bounceIn: {
          '0%': { transform: 'scale(0.3)', opacity: '0' },
          '50%': { transform: 'scale(1.05)' },
          '70%': { transform: 'scale(0.9)' },
          '100%': { transform: 'scale(1)', opacity: '1' }
        },
        slideUp: {
          'from': { transform: 'translateY(20px)', opacity: '0' },
          'to': { transform: 'translateY(0)', opacity: '1' }
        },
        fadeIn: {
          'from': { opacity: '0' },
          'to': { opacity: '1' }
        },
        pulseOrange: {
          '0%, 100%': { opacity: '1' },
          '50%': { opacity: '.5' }
        }
      },
      
      backgroundImage: {
        'gradient-cear': 'linear-gradient(135deg, #f97316 0%, #ea580c 100%)',
        'gradient-cear-reverse': 'linear-gradient(135deg, #ea580c 0%, #f97316 100%)',
        'gradient-radial-cear': 'radial-gradient(ellipse at center, #f97316 0%, #ea580c 100%)'
      },
      
      borderRadius: {
        'cear': '0.5rem'
      },
      
      spacing: {
        '18': '4.5rem',
        '88': '22rem',
        '128': '32rem'
      }
    }
  },
  
  plugins: [
    function({ addUtilities, addComponents, theme }) {
      // Utilities customizadas CEAR
      const newUtilities = {
        '.bg-gradient-cear': {
          background: 'linear-gradient(135deg, #f97316 0%, #ea580c 100%)'
        },
        '.bg-gradient-cear-reverse': {
          background: 'linear-gradient(135deg, #ea580c 0%, #f97316 100%)'
        },
        '.bg-gradient-radial-cear': {
          background: 'radial-gradient(ellipse at center, #f97316 0%, #ea580c 100%)'
        },
        '.text-gradient-cear': {
          background: 'linear-gradient(135deg, #f97316 0%, #ea580c 100%)',
          '-webkit-background-clip': 'text',
          '-webkit-text-fill-color': 'transparent',
          'background-clip': 'text'
        },
        '.border-gradient-cear': {
          border: '2px solid transparent',
          'background-clip': 'padding-box',
          'background-image': 'linear-gradient(135deg, #f97316 0%, #ea580c 100%)'
        }
      }
      
      // Componentes CEAR reutilizáveis
      const newComponents = {
        '.btn-cear': {
          '@apply bg-gradient-cear hover:from-primary-600 hover:to-primary-700': {},
          '@apply text-white font-medium px-6 py-3 rounded-lg': {},
          '@apply transition-all duration-200 shadow-orange': {},
          '@apply focus:ring-2 focus:ring-primary-500 focus:ring-opacity-50': {},
          '@apply active:scale-95': {}
        },
        '.btn-cear-outline': {
          '@apply border-2 border-primary-500 text-primary-600': {},
          '@apply hover:bg-primary-50 focus:bg-primary-50': {},
          '@apply font-medium px-6 py-3 rounded-lg': {},
          '@apply transition-all duration-200': {},
          '@apply focus:ring-2 focus:ring-primary-500 focus:ring-opacity-50': {}
        },
        '.card-cear': {
          '@apply bg-white rounded-lg shadow-orange-md border border-secondary-100': {},
          '@apply hover:shadow-orange-lg transition-all duration-300': {},
          '@apply p-6': {}
        },
        '.input-cear': {
          '@apply w-full px-4 py-3 border-2 border-secondary-300 rounded-lg': {},
          '@apply focus:ring-2 focus:ring-primary-500 focus:border-primary-500': {},
          '@apply transition-all duration-200': {},
          '@apply placeholder-secondary-400': {}
        },
        '.badge-cear': {
          '@apply inline-flex items-center px-3 py-1 rounded-full text-sm font-medium': {},
          '@apply bg-primary-100 text-primary-800': {}
        },
        '.navbar-cear': {
          '@apply bg-white shadow-orange border-b border-secondary-200': {}
        },
        '.footer-cear': {
          '@apply bg-secondary-900 text-secondary-100': {}
        }
      }
      
      addUtilities(newUtilities)
      addComponents(newComponents)
    }
  ]
} 