# App de Memorias Digitales - Documentación Técnica

## Stack Tecnológico

### Backend
- Go
- PostgreSQL
- AWS S3 (almacenamiento de imágenes)

### Frontend
- SvelteKit
- TailwindCSS

## Análisis del Proyecto

### Problemática y Usuarios
- **Problema:** Organizar y mantener recuerdos digitales de manera personal
- **Usuarios objetivo:** Mi novia y sus amigas viajeras

### Objetivos
- **Funcional:** Crear una plataforma para gestionar recuerdos visuales
- **Técnico:** Desarrollar CRUDs en Go y mejorar lógica de programación

### Priorización de Funcionalidades

#### MVP (Esenciales)
- Autenticación de usuarios
- Subida de múltiples fotos
- Creación de historias/descripciones
- Organización por fechas
- Etiquetado básico de ubicaciones

#### Versión 2.0
- Sistema de amigos
- Tags avanzados
- Búsqueda multi-parámetro
- Versión demo con localStorage

### Limitaciones
- Curva de aprendizaje de Go
- Restricciones de tiempo
- Mejora de habilidades con SvelteKit

## User Stories

Como usuario quiero:

1. Subir múltiples fotos de un evento/fecha
2. Escribir historias/descripciones de los momentos
3. Categorizar mis recuerdos
4. Buscar mis recuerdos fácilmente
5. Etiquetar ubicaciones de los momentos

## Criterios de Éxito

### Técnicos
- CRUD funcional completo
- Sistema eficiente de subida de fotos
- Búsquedas rápidas y precisas

### Usuario
- Interfaz intuitiva y amigable
- Tiempos de carga optimizados
- Organización simple e intuitiva de recuerdos

## Especificaciones Técnicas

### Modelo de Datos
- **Tipos de datos:** Fotos, texto, fechas, categorías, ubicaciones, usuarios
- **Relaciones:**
 - 1:N entre recuerdos y fotos
 - 1:N entre recuerdos y categorías
 - 1:1 entre recuerdos y ubicaciones
 - 1:1 entre recuerdos y propietario

### Seguridad
- **Autenticación:** JWT + cookies de sesión
- **Almacenamiento seguro:** AWS S3 para fotos
- **Protección de datos:** Encriptación + HTTPS
- **Opciones adicionales:** OAuth2 para integración futura

## Wireframes Principales

### Pantallas Esenciales
1. **Login/Registro**
  - Formulario minimalista
  - Opciones de recuperación de contraseña

2. **Dashboard Principal**
  - Vista de recuerdos recientes
  - Acceso rápido a categorías
  - Barra de búsqueda destacada

3. **Subida de Contenido**
  - Área de arrastrar y soltar para fotos
  - Editor de texto para descripciones
  - Selector de categorías y ubicación

4. **Vista de Recuerdos**
  - Galería visual con opciones de filtrado
  - Visualización de historia/descripción
  - Metadatos asociados (fecha, ubicación, etc.)

5. **Sistema de Categorías**
  - Gestión de etiquetas personalizadas
  - Visualización de recuerdos por categoría

### Flujo de Navegación
- Diseño intuitivo con navegación jerárquica
- Breadcrumbs para ubicación contextual
- Accesos rápidos a funciones principales